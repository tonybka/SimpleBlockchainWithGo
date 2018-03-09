package main

import (
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt16
)

const targetBits = 16

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

//Run
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int

	fmt.Printf("Mining a new block with POW")
	var hash [32]byte
	nonce := 0

	return nonce, hash[:]
}
