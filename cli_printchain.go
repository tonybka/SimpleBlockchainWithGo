package main

func (cli *CLI) printBlockchain(nodeID string) {
	newBc := NewBlockchain(nodeID)
	defer newBc.db.Close()
}
