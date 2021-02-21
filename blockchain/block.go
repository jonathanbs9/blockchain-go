package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
	Nonce int
}

// Creamos el bloque
func CreateBlock( data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Agregamos un bloque a la cadena de bloques
func (chain *BlockChain) AddBlock(data string){
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Función Genesis => Crea el primer bloque
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Función que inicializa la cadena de bloques (precisa del genesis)
func InitBlockChain() *BlockChain{
	return &BlockChain{[]*Block{Genesis()}}
}