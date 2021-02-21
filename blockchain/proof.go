package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Tomamos le informacion (Data) del bloque

// Creamos un contador que arranca en 0

// Creamos un Hash de la información + el contador

// Chequeamos el Hash para ver si concuerda con el conjunto de requerimientos

// Requerimientos:
// Los primeros bytes deben contener ceros

// Cuando mayor es el difficulty mas complejo resultará
const Difficulty = 3

func (b *Block) DeriveHash(){
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash:= sha256.Sum256(info)
	b.Hash = hash[:]
}

type ProofOfWork struct {
	Block *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target:= big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData (nonce int) []byte  {
	data := bytes.Join(
			[][]byte{
				pow.Block.PrevHash,
				pow.Block.Data,
				ToHex(int64(nonce)),
				ToHex(int64(Difficulty)),
			},
			[]byte{},
		)
	return data
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic("Error al convertir num => ",err)
	}
	return buff.Bytes()
}

func (pow *ProofOfWork) Run() (int, []byte){
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64{
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r %x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)
	// Tomamos la data y la convertimos en hash
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
