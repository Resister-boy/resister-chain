package main

import (
	"fmt"
	"./blockchain/blockchain.go"
)

func main() {
	chain := blockchain.getBlockchain()	// Blockchain Generate Func
	chain.AddBlock("Second Block") // Generate Block
	chain.AddBlock("Third Block")
	chain.AddBlock("Forth Block")
	chain.AddBlock("Fifth Block")
	for _, block := range chain.AllBlocks() { // Block Data 출력
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %s\n", block.Hash)
		fmt.Printf("Previous Hash : %s\n", block.Hash)
	}
}