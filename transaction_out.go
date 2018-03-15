package main

//TXoutput represent a transaction output
type TXoutput struct {
	Value      int
	PubKeyHash []byte
}

//LockSigns locks signs the out put
func (txOut *TXoutput) LockSigns(address []byte) {

}
