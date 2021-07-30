package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/jonathanbs9/blockchain-go/blockchain"
)

type CommandLine struct {
	blockchain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add - block BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - Prints the blocks in the chain")
}

// func validateArgs = allow us to validate any arg that we pass through the command line
func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		// Exits the application by shutting down de GORoutine
		runtime.Goexit()
	}
}

// func addBlock
func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Bloque agregado!!! ")
}

// func printChain
func (cli *CommandLine) printChain() {
	// Converts the blockchain struct into an Iterator Struct
	iter := cli.blockchain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Hash previo => %x \n", block.PrevHash)
		fmt.Printf("Data en el bloque => %x \n", block.Data)
		fmt.Printf("Hash => %x \n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("Pow =>  %s \n \n", strconv.FormatBool(pow.Validate()))

		if len(block.PrevHash) == 0 {
			break
		}
	}

}

// func run
func (cli *CommandLine) run() {
	cli.validateArgs()
	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

}

func main() {
	defer os.Exit(0)
	//chain := blockchain.InitBlockChain()

	/*chain.AddBlock("Primer bloque despues del Genesis")
	chain.AddBlock("Segundo bloque despues del Genesis")
	chain.AddBlock("Tercer bloque despues del Genesis")
	chain.AddBlock("Cuarto bloque despues del Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Hash previo => %x \n", block.PrevHash)
		fmt.Printf("Data en el bloque => %x \n", block.Data)
		fmt.Printf("Hash => %x \n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("Pow =>  %s \n \n", strconv.FormatBool(pow.Validate()))

	}*/
	chain := blockchain.InitBlockChain()
	// Only executes if the go channel is able to exit properly
	defer chain.Database.Close()
	cli := CommandLine{chain}
	cli.run()
}
