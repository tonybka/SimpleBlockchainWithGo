package main

import (
	"bytes"
)

//TXinput resperent a transaction input
type TXinput struct {
	TxID      []byte
	VarOut    int
	Signature []byte
	PubKey    []byte
}

//UsesKey checks whether address initiated this transaction
func (txIn *TXinput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(txIn.PubKey)
	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
