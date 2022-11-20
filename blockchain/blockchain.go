package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Hash string
	Data string
	PrevHash string
}

type blockchain struct {
	blocks []*Block
}

func (b *Block) createHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(getBlockchain().blocks)
	if (totalBlocks == 0) {
		return ""
	}
	return getBlockchain().blocks[totalBlocks - 1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block(data, "", getLastHash())
	newBlock.createHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func getBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}