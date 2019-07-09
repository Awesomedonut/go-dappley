// Copyright (C) 2018 go-dappley authors
//
// This file is part of the go-dappley library.
//
// the go-dappley library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-dappley library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dappley library.  If not, see <http://www.gnu.org/licenses/>.
//

package network

import (
	"bufio"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	logger "github.com/sirupsen/logrus"
)

const (
	highPriorityChLength   = 1024 * 4
	normalPriorityChLength = 1024 * 4
	WriteChTotalLength     = highPriorityChLength + normalPriorityChLength
)

const (
	HighPriorityCommand = iota
	NormalPriorityCommand
)

type Stream struct {
	peerID                peer.ID
	remoteAddr            multiaddr.Multiaddr
	stream                network.Stream
	rawByteRead           []byte
	highPriorityWriteCh   chan []byte
	normalPriorityWriteCh chan []byte
	msgNotifyCh           chan bool
	quitRdCh              chan bool
	quitWrCh              chan bool
}

func NewStream(s network.Stream) *Stream {
	return &Stream{
		s.Conn().RemotePeer(),
		s.Conn().RemoteMultiaddr(),
		s,
		[]byte{},
		make(chan []byte, highPriorityChLength),
		make(chan []byte, normalPriorityChLength),
		make(chan bool, WriteChTotalLength),
		make(chan bool, 1), //two channels to stop
		make(chan bool, 1),
	}
}

func (s *Stream) Start(quitCh chan<- *Stream, dispatch chan *streamMsg) {
	logger.Warn("Stream: Start new stream")
	rw := bufio.NewReadWriter(bufio.NewReader(s.stream), bufio.NewWriter(s.stream))
	s.startLoop(rw, quitCh, dispatch)
}

func (s *Stream) StopStream(err error) {
	logger.WithFields(logger.Fields{
		"peer_address": s.remoteAddr,
		"pid":          s.peerID,
		"error":        err,
	}).Info("Stream: is terminated!!")
	s.quitRdCh <- true
	s.quitWrCh <- true
	s.stream.Close()
}

func (s *Stream) Send(data []byte, priority int) {
	defer func() {
		if p := recover(); p != nil {
			logger.WithFields(logger.Fields{
				"peer_address": s.remoteAddr,
				"pid":          s.peerID,
				"error":        p,
			}).Info("Stream: data channel closed.")
		}
	}()

	switch priority {
	case HighPriorityCommand:
		select {
		case s.highPriorityWriteCh <- data:
		default:
			logger.WithFields(logger.Fields{
				"dataCh_len": len(s.highPriorityWriteCh),
			}).Warn("Stream: High priority message channel full!")
			return
		}
	case NormalPriorityCommand:
		select {
		case s.normalPriorityWriteCh <- data:
		default:
			logger.WithFields(logger.Fields{
				"dataCh_len": len(s.normalPriorityWriteCh),
			}).Warn("Stream: normal priority message channel full!")
			return
		}
	default:
		logger.WithFields(logger.Fields{
			"priority": priority,
		}).Warn("Stream: priority is invalid!")
		return
	}

	select {
	case s.msgNotifyCh <- true:
	default:
		logger.WithFields(logger.Fields{
			"dataCh_len": len(s.msgNotifyCh),
		}).Warn("Stream: message notification channel full!")
	}

}

func (s *Stream) startLoop(rw *bufio.ReadWriter, quitCh chan<- *Stream, dispatch chan *streamMsg) {
	go s.readLoop(rw, quitCh, dispatch)
	go s.writeLoop(rw)
}

func (s *Stream) read(rw *bufio.ReadWriter, dispatch chan *streamMsg) {
	buffer := make([]byte, 1024)
	var err error

	n, err := rw.Read(buffer)
	if err != nil {
		logger.WithError(err).WithFields(logger.Fields{
			"num_of_byte_read": n,
		}).Warn("Stream: Read failed")
		s.StopStream(err)
	}

	s.rawByteRead = append(s.rawByteRead, buffer[:n]...)

	for {
		packet, err := ExtractDappPacketFromRawBytes(s.rawByteRead)

		if err != nil {
			if err == ErrLengthTooShort {
				return
			} else {
				s.StopStream(err)
			}
		}

		dm := ParseDappMsgFromRawBytes(packet.GetData())
		select {
		case dispatch <- newMsg(dm, s.peerID):
		default:
			logger.WithFields(logger.Fields{
				"dispatchCh_len": len(dispatch),
			}).Warn("Stream: command dispatch channel full!")
			return
		}
		s.rawByteRead = s.rawByteRead[packet.GetLength():]
	}

}

func (s *Stream) readLoop(rw *bufio.ReadWriter, quitCh chan<- *Stream, dispatch chan *streamMsg) {
	for {
		select {
		case <-s.quitRdCh:
			quitCh <- s
			logger.Debug("Stream: read loop is terminated!")
			return
		default:
			s.read(rw, dispatch)
		}
	}
}

func (s *Stream) writeLoop(rw *bufio.ReadWriter) error {

	for {
		select {
		case <-s.quitWrCh:
			// Fix bug when send data to peer simultaneous with close stream,
			// and send will hang because highPriorityWriteCh is full.
			close(s.highPriorityWriteCh)
			close(s.normalPriorityWriteCh)
			close(s.msgNotifyCh)
			logger.Debug("Stream: write loop is terminated!")
			return nil
		case <-s.msgNotifyCh:
			select {
			case data := <-s.highPriorityWriteCh:
				packet := ConstructDappPacketFromData(data)
				n, err := s.stream.Write(packet.GetRawBytes())
				if err != nil {
					logger.WithError(err).WithFields(logger.Fields{
						"num_of_bytes_sent": n,
						"orig_data_size":    packet.GetLength(),
					}).Warn("Stream: Send message through high priority channel failed!")
				}
				continue
			default:
			}
			select {
			case data := <-s.normalPriorityWriteCh:
				packet := ConstructDappPacketFromData(data)
				n, err := s.stream.Write(packet.GetRawBytes())
				if err != nil {
					logger.WithError(err).WithFields(logger.Fields{
						"num_of_bytes_sent": n,
						"orig_data_size":    packet.GetLength(),
					}).Warn("Stream: Send message through normal priority channel failed!")
				}
				continue
			default:
			}
		}

	}
	return nil
}
