package main

/**
- Output is indivisible, spent as a whole in new transaction
- If its value greater than required, a change is generated and sent back to sender
**/
//TXoutput represent a transaction output
type TXoutput struct {
	Value      int
	PubKeyHash []byte
}

//LockSigns locks signs the out put
func (txOut *TXoutput) LockSigns(address []byte) {

}
