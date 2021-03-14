package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

// Ahora con Badger va a almacenar el ultimo hash del ultimo bloque de cadena
// Y un puntero al badgerDB
type BlockChain struct {
	//Blocks []*Block
	LastHash []byte
	Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{chain.LastHash, chain.Database}
	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block
	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		encodedBlock, err := item.Value()
		block = Deserialize(encodedBlock)

		return err
	})
	Handle(err)
	iter.CurrentHash = block.PrevHash
	return block
}

// Función que inicializa la cadena de bloques (precisa del genesis)
func InitBlockChain() *BlockChain {
	//return &BlockChain{[]*Block{Genesis()}}
	var lastHash []byte
	opts := badger.DefaultOptions
	// Donde la base de datos va a almacenar los keys y metedata
	opts.Dir = dbPath
	// Donde la base de datos va a almacenar todos los valores
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		/* Chequemaos en la bd si el blockchain tiene algo almacenado.
		/  Si ya tiene, creamos una instancia nueva  de blockchain en memoria  y tenemos el lastHash del blockchain en el disk DB
		/  y la pusheamos en la instancia de blockchain.
		*/
		/* Si no hay ninguna instancia de Blockhain en DB, creamos un blocke Genesis y lo almacenamos en la BD. Guardamos el hash del  bloque Genesis
		   como el lastHash en BD y creamos una nueva instancia de  blockchain apuntando al bloque Genesis
		*/
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No se encontró Blockchain")
			genesis := Genesis()
			fmt.Println("Geneis probado!")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			lastHash, err = item.Value()
			return err
		}
	})

	Handle(err)
	// Creamos un nuevo blockchain en memoria
	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

// Agregamos un bloque a la cadena de bloques
func (chain *BlockChain) AddBlock(data string) {
	//prevBlock := chain.Blocks[len(chain.Blocks)-1]
	//new := CreateBlock(data, prevBlock.Hash)
	//chain.Blocks = append(chain.Blocks, new)
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()

		return err
	})
	Handle(err)

	newBlock := CreateBlock(data, lastHash)
	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})
	Handle(err)
}
