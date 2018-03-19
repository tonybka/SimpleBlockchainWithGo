package main

import (
	"flag"
	"fmt"
	"os"
)

//CLI responsible for processing blockchain COMMAND LINE arguments
type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("  createwallet - Generates a new key-pair and saves it into the wallet file")
	fmt.Println("  getbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("  listaddresses - Lists all addresses from the wallet file")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
	fmt.Println("  reindexutxo - Rebuilds the UTXO set")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT -mine - Send AMOUNT of coins from FROM address to TO. Mine on the same node, when -mine is set.")
	fmt.Println("  startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining")
}

func (cli *CLI) validateArgs() {
	//if number of args is less than 2, default the fisrt arg is path of current excutable file
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

//Run parses command arguments and process command
func (cli *CLI) Run() {
	cli.validateArgs()
	nodeID := os.Getenv("NODE_ID")

	if nodeID == "" {
		fmt.Printf("NODE_ID env value is not set!")
		os.Exit(1)
	}

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)

	switch os.Args[1] {
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		break
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		break
	default:
		cli.printUsage()
		os.Exit(1)
	}
}
