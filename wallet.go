package main

import (
	"crypto/sha256"
	"log"

	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)
const addressChecksumLen = 4

//Wallet store private key and public key
type Wallet struct {
}

//HashPubKey hash the public key
func HashPubKey(pubkey []byte) []byte {
	pubSHA256 := sha256.Sum256(pubkey)
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(pubSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

func checksum(payLoad []byte) []byte {
	firstSHA256 := sha256.Sum256(payLoad)
	secondSHA256 := sha256.Sum256(firstSHA256[:])
	return secondSHA256[:addressChecksumLen]
}
