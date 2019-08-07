package core

import (
	"encoding/hex"

	"github.com/dappley/go-dappley/core/account"
	logger "github.com/sirupsen/logrus"
)

//FindVinUtxosInUtxoPool Find the transaction in a utxo pool. Returns true only if all Vins are found in the utxo pool
func FindVinUtxosInUtxoPool(utxoPool UTXOIndex, tx Transaction) ([]*UTXO, error) {
	var res []*UTXO
	for _, vin := range tx.Vin {
		pubKeyHash, err := account.NewUserPubKeyHash(vin.PubKey)
		if err != nil {
			return nil, ErrNewUserPubKeyHash
		}
		utxo := utxoPool.FindUTXOByVin([]byte(pubKeyHash), vin.Txid, vin.Vout)
		if utxo == nil {
			logger.WithFields(logger.Fields{
				"txid":      hex.EncodeToString(tx.ID),
				"vin_id":    hex.EncodeToString(vin.Txid),
				"vin_index": vin.Vout,
			}).Warn("Transaction: Can not find vin")
			return nil, ErrTXInputNotFound
		}
		res = append(res, utxo)
	}
	return res, nil
}