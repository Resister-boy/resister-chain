package main

import (
	"fmt"
	"crypto/sha256"
)

type block struct {
	data string
	hash string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) getLastBlockHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks) - 1].hash // 가장 최근에 생성된 블록의 Hash
	}
	return ""
}

func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", b.getLastBlockHash()} // Insert 이전 블록 Hash
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash)) // Create 현재 블록 Hash
	newBlock.hash = fmt.Sprintf("%x", hash)
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data : %s\n", block.data)
		fmt.Printf("Hash : %s\n", block.hash)
		fmt.Printf("PrevHash : %s\n", block.prevHash)
	}
}

func main() {
	chain := blockchain{}
	chain.addBlock("Genesis Block") // data parameter
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
}