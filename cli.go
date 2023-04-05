package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"strings"

	"os"
)

// CLI responsible for processing command line arguments
type CLI struct {
	bc *Blockchain
}

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
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments and processes commands
// func (cli *CLI) Run() {
// 	cli.validateArgs()

// 	nodeID := os.Getenv("NODE_ID")
// 	if nodeID == "" {
// 		fmt.Printf("NODE_ID env. var is not set!")
// 		os.Exit(1)
// 	}

// 	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
// 	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
// 	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
// 	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)
// 	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
// 	reindexUTXOCmd := flag.NewFlagSet("reindexutxo", flag.ExitOnError)
// 	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
// 	startNodeCmd := flag.NewFlagSet("startnode", flag.ExitOnError)

// 	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
// 	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
// 	sendFrom := sendCmd.String("from", "", "Source wallet address")
// 	sendTo := sendCmd.String("to", "", "Destination wallet address")
// 	sendAmount := sendCmd.Int("amount", 0, "Amount to send")
// 	sendMine := sendCmd.Bool("mine", false, "Mine immediately on the same node")
// 	startNodeMiner := startNodeCmd.String("miner", "", "Enable mining mode and send reward to ADDRESS")

// 	switch os.Args[1] {
// 	case "getbalance":
// 		err := getBalanceCmd.Parse(os.Args[2:])
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	case "createblockchain":
// 		err := createBlockchainCmd.Parse(os.Args[2:])
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	case "createwallet":
// 		err := createWalletCmd.Parse(os.Args[2:])
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	case "listaddresses":
// 		err := listAddressesCmd.Parse(os.Args[2:])
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	case "printchain":
// 		err := printChainCmd.Parse(os.Args[2:])
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	case "reindexutxo":
// 		err := reindexUTXOCmd.Parse(os.Args[2:])
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	case "send":
// 		err := sendCmd.Parse(os.Args[2:])
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	case "startnode":
// 		err := startNodeCmd.Parse(os.Args[2:])
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	default:
// 		cli.printUsage()
// 		os.Exit(1)
// 	}

// 	if getBalanceCmd.Parsed() {
// 		if *getBalanceAddress == "" {
// 			getBalanceCmd.Usage()
// 			os.Exit(1)
// 		}
// 		cli.getBalance(*getBalanceAddress, nodeID)
// 	}

// 	if createBlockchainCmd.Parsed() {
// 		if *createBlockchainAddress == "" {
// 			createBlockchainCmd.Usage()
// 			os.Exit(1)
// 		}
// 		cli.createBlockchain(*createBlockchainAddress, nodeID)
// 	}

// 	if createWalletCmd.Parsed() {
// 		cli.createWallet(nodeID)
// 	}

// 	if listAddressesCmd.Parsed() {
// 		cli.listAddresses(nodeID)
// 	}

// 	if printChainCmd.Parsed() {
// 		cli.printChain(nodeID)
// 	}

// 	if reindexUTXOCmd.Parsed() {
// 		cli.reindexUTXO(nodeID)
// 	}

// 	if sendCmd.Parsed() {
// 		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
// 			sendCmd.Usage()
// 			os.Exit(1)
// 		}

// 		cli.send(*sendFrom, *sendTo, *sendAmount, nodeID, *sendMine)
// 	}

// 	if startNodeCmd.Parsed() {
// 		nodeID := os.Getenv("NODE_ID")
// 		if nodeID == "" {
// 			startNodeCmd.Usage()
// 			os.Exit(1)
// 		}
// 		cli.startNode(nodeID, *startNodeMiner)
// 	}

// }

//changed Run()
func (cli *CLI) Run() {
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		fmt.Printf("NODE_ID env. var is not set!")
		os.Exit(1)
	}

	cli.printUsage()
	for {
		fmt.Print("\nEnter a command: ")
		reader := bufio.NewReader(os.Stdin)
		inputString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(inputString) == 0 {
			continue
		}

		inputString = inputString[:len(inputString)-1]
		commandArgs := strings.Fields(inputString)

		//test
		// fmt.Println(commandArgs[0])

		switch commandArgs[0] {
		case "getbalance":
			getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
			address := getBalanceCmd.String("address", "", "The address to get balance for")
			if len(commandArgs) != 3 {
				getBalanceCmd.Usage()
				continue
			}
			getBalanceCmd.Parse(commandArgs[1:])
			cli.getBalance(*address, nodeID)
		case "createblockchain":
			createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
			address := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
			// if len(commandArgs) != 3 {
			// 	createBlockchainCmd.Usage()
			// 	continue
			// }
			createBlockchainCmd.Parse(commandArgs[1:])
			cli.createBlockchain(*address, nodeID)
		case "createwallet":
			cli.createWallet(nodeID)
		case "listaddresses":
			cli.listAddresses(nodeID)
		case "printchain":
			cli.printChain(nodeID)
		case "reindexutxo":
			cli.reindexUTXO(nodeID)
		case "send":
			sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
			from := sendCmd.String("from", "", "Source wallet address")
			to := sendCmd.String("to", "", "Destination wallet address")
			amount := sendCmd.Int("amount", 0, "Amount to send")
			mine := sendCmd.Bool("mine", false, "Mine immediately on the same node")
			if len(commandArgs) != 8 && len(commandArgs) != 7 {
				sendCmd.Usage()
				continue
			}
			sendCmd.Parse(commandArgs[1:])
			cli.send(*from, *to, *amount, nodeID, *mine)
		case "startnode":
			startNodeCmd := flag.NewFlagSet("startnode", flag.ExitOnError)
			minerAddress := startNodeCmd.String("miner", "", "Enable mining mode and send reward to ADDRESS")
			// if len(commandArgs) != 2 {
			// 	startNodeCmd.Usage()
			// 	continue
			// }
			startNodeCmd.Parse(commandArgs[1:])
			cli.startNode(nodeID, *minerAddress)
		case "usage":
			cli.printUsage()
		case "exit":
			return
		default:
			fmt.Printf("Unknown command: %s\n", commandArgs[0])
			continue
		}
	}
}
