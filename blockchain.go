package main

import (
	"bytes"
	"errors"
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
	var tip []byte
	cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
	genesisBlock := NewGenesisBlock(cbtx)
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		// create bucket to store block data
		bucket, err := tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			log.Panic(err)
		}

		// map data with corresponding keys
		err = bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = bucket.Put([]byte("l"), genesisBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		tip = genesisBlock.Hash
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBc := Blockchain{tip, db}

	return &newBc
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
func (bc *Blockchain) FindTransaction(txID []byte) (Transaction, error) {
	bci := bc.Iterator()
	for {
		//iterate over each block on blockchain
		block := bci.NextBLock()

		for _, tx := range block.Transactions {
			if bytes.Compare(tx.ID, txID) == 0 {
				return *tx, nil
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return Transaction{}, errors.New("Transaction not found")
}

//AddBlock concatenate block into the blockchain
func (bc *Blockchain) AddBlock(block *Block) {
	err := bc.db.Update(func(tx *bolt.Tx) error {
		// Get bucket where saving blockchain data
		bucket := tx.Bucket([]byte(blocksBucket))
		blockInDb := bucket.Get(block.Hash)

		// If this block already existed on this current blockchain
		if blockInDb != nil {
			return nil
		}

		blockData := block.Serialize()
		err := bucket.Put(block.Hash, blockData)

		if err != nil {
			log.Panic(err)
		}
		lastHash := bucket.Get([]byte("l"))
		lastBlockData := bucket.Get(lastHash)
		lastBlock := DeserializeBlock(lastBlockData)

		if block.Height > lastBlock.Height {
			err := bucket.Put([]byte("l"), block.Hash)
			if err != nil {
				log.Panic(err)
			}
			bc.tip = block.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
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

//Iterator return a BlockchainIterator
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bcIterator := &BlockchainIterator{bc.tip, bc.db}
	return bcIterator
}
