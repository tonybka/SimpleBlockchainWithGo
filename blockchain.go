package main

const dbFile = "blockchain_%s.db"

//Blockchain implement interaction with a DB
type Blockchain struct {
}

//AddBlock concatenate block into the blockchain
func (bc *Blockchain) AddBlock(block *Block) {
	// Get previous block hash
}

//NewGenesisBlock create first block of blockchain
func NewGenesisBlock() {

}
