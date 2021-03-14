package main

import (
	"blockchain-go/blockchain"
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

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Bloque agregado!!! ")
}

func (cli *CommandLine) printChain() {
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

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Primer bloque despues del Genesis")
	chain.AddBlock("Segundo bloque despues del Genesis")
	chain.AddBlock("Tercer bloque despues del Genesis")
	chain.AddBlock("Cuarto bloque despues del Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Hash previo => %x \n", block.PrevHash)
		fmt.Printf("Data en el bloque => %x \n", block.Data)
		fmt.Printf("Hash => %x \n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("Pow =>  %s \n \n", strconv.FormatBool(pow.Validate()))

	}

}
