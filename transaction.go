package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"log"
)

const subsidy = 10

/**
 - There are outputs that not linked to inputs
 - Inputs can reference outputs from multiple transactions
 - An input must reference an output
**/
// Transaction represents a Bitcoin transaction
type Transaction struct {
	ID     []byte
	Vin    []TXinput
	Vout   []TXoutput
	numIn  int
	numOut int
}

// Serialize returns a serialized Transaction
func (tx Transaction) Serialize() []byte {
	var encoded bytes.Buffer
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return encoded.Bytes()
}

//DeserializeTransaction deserializes a transaction
func DeserializeTransaction(data []byte) Transaction {

}

// IsCoinbase checks whether the transaction is coinbase
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].TxID) == 0 && tx.Vin[0].VarOut == -1
}

//Verify verifies signatures of transaction input
func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}

	for _, vin := range tx.Vin {
		if prevTXs[hex.EncodeToString(vin.TxID)].ID == nil {
			log.Panic("ERROR: Previous transaction is not correct")
		}
	}

}

//TrimmedCopy create trimed copy of transaction to be used in signing
func (tx *Transaction) TrimmedCopy() Transaction {

}

//NewCoinbaseTX create new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {

}

//NewUTXOTransaction creates a new transaction
func NewUTXOTransaction(wallet *Wallet, to string, amount int) *Transaction {

}
