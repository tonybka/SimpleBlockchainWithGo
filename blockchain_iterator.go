package main

import (
	"log"

	"github.com/boltdb/bolt"
)

//BlockchainIterator used to iterate over blocks of blockchain
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

//NextBLock return next block starting from the tip
func (bci *BlockchainIterator) NextBLock() *Block {
	var block *Block

	err := bci.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(blocksBucket))
		encodedBlock := bk.Get(bci.currentHash)
		block := DeserializeBlock(encodedBlock)
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	bci.currentHash = block.PrevBlockHash
	return block
}
