package main

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain_%s.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/FebJan/2018: Start new blockchain"

//Blockchain implement interaction with a DB
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

//CreateBlockchain create new blockchain DB
func CreateBlockchain(address, nodeID string) *Blockchain {
	dbFile := fmt.Sprintf(dbFile, nodeID)
	if isDbExisted(dbFile) {
		fmt.Println("Blockchain already exists.")
		os.Exit(1)
	}
}

//VerifyTransaction verifies transaction input signatures
func (bc *Blockchain) VerifyTransaction(tx *Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}
	prevTXs := make(map[string]Transaction)
	for _, vin := range tx.Vin {
		prevTX, err := bc.FindTransaction(vin.TxID)

	}

	return tx.Verify(prevTXs)
}

//FindTransaction find a transaction by its ID
func (bc *Blockchain) FindTransaction(txId []byte) (Transaction, error) {

}

//AddBlock concatenate block into the blockchain
func (bc *Blockchain) AddBlock(block *Block) {
	// Get previous block hash
}

//NewBlockchain create new blockchain with genesis block
func NewBlockchain(nodeID string) *Blockchain {
	dbFile := fmt.Sprintf(dbFile, nodeID)
	if isDbExisted(dbFile) == false {
		fmt.Println("No existing blockchain found. Please create new one.")
		os.Exit(1)
	}
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		block := tx.Bucket([]byte(blocksBucket))
		//the last block file number used
		tip = block.Get([]byte("l"))
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	bc := Blockchain{tip, db}

	return &bc
}

//isDbExisted checks blockchain database was created or not
func isDbExisted(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}
