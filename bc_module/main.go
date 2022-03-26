package main

import (
	"bytes"
	"crypto/sha256"
)

func main() {

}

type simple struct {
	Set1 []byte // uint8
	Set2 []byte
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

// hash data given to block
func (b *Block) SetBlockHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{}) // split into 2d-array then flatten into 1d
	hash := sha256.Sum256((info))                              // apply hash
	b.Hash = hash[:]                                           // store hashed data + prevhash
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.SetBlockHash()

	return block
}
