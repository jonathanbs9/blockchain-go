package main

import (
	"blockchain-go/blockchain"
	"fmt"
	"strconv"
)

/*type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// Función para Crear un Block
func CreateBlock(data string, prevHash []byte) *Block{
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// Función para Agregar un Block
func (chain *BlockChain) AddBlock(data string){
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Función para crear un Block inicializador
func Genesis() *Block{
	return CreateBlock("Genesis", []byte{})
}

// Funcion para inicializar el BlockChain
func InitBlockChain() *BlockChain{
	return &BlockChain{[]*Block{Genesis()}}
}*/

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Primer bloque despues del Genesis")
	chain.AddBlock("Segundo bloque despues del Genesis")
	chain.AddBlock("Tercer bloque despues del Genesis")
	chain.AddBlock("Cuarto bloque despues del Genesis")

	for   _, block := range chain.Blocks{
		fmt.Printf("Hash previo => %x \n", block.PrevHash)
		fmt.Printf("Data en el bloque => %x \n", block.Data)
		fmt.Printf("Hash => %x \n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("Pow =>  %s \n \n", strconv.FormatBool(pow.Validate()))

	}

}
