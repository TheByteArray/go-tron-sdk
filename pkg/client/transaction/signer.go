package transaction

import (
	"crypto/ecdsa"
	"crypto/sha256"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/TheByteArray/go-tron-sdk/pkg/proto/core"
	proto "google.golang.org/protobuf/proto"
)

type SignerImpl int

const (
	Software SignerImpl = iota
	Ledger
)

func SignTransaction(tx *core.Transaction, signer *btcec.PrivateKey) (*core.Transaction, error) {
	return SignTransactionECDSA(tx, signer.ToECDSA())
}

func SignTransactionECDSA(tx *core.Transaction, signer *ecdsa.PrivateKey) (*core.Transaction, error) {
	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return nil, err
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	signature, err := crypto.Sign(hash, signer)
	if err != nil {
		return nil, err
	}
	tx.Signature = append(tx.Signature, signature)

	return tx, nil
}
