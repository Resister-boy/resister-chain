package blockchain

import (
	"sync",
	"fmt"
)

type block struct { // Block Components
	Data string
	Hash string
	PrevHash string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func (b *block) createHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash)) // Current Block Hash Generate
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string { // return Lash Block Hash
	totalBlocks := len(getBlockchain().blocks)
	if (totalBlocks == 0) {
		return "" // if Genesis Block return null
	}
	return getBlockchain().blocks[totalBlocks - 1].Hash
	// ifn Genesis Block return Previous Block Hash
}

func createBlock(data string) *block {
	newBlock := block(data, "", getLastHash)
	newBlock.createHash() // Create Hash Func
	return newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func getBlockchain() *blockchain {
	if b == nil { // if문은 한 번만 실행, 한 번 실행되면 b는 nil이 될 수 없기 때문
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() {}*block {
	return b.block
}