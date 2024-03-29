package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PrevHash     string
	Hash         string
	Nonce        int
}

type Blockchain struct {
	Blocks []*Block
}

func NewGenesisBlock() *Block {
	return NewBlock(0, "Genesis Block", "")
}

func NewBlock(index int, data, prevHash string) *Block {
	block := &Block{Index: index, Timestamp: time.Now().String(), Data: data, PrevHash: prevHash}
	block.Hash = block.CalculateHash()
	block.MineBlock(difficulty)
	return block
}

func (b *Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + strconv.Itoa(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

const difficulty = 3

func (b *Block) MineBlock(difficulty int) {
	prefix := strings.Repeat("0", difficulty)
	for !strings.HasPrefix(b.Hash, prefix) {
		b.Nonce++
		b.Hash = b.CalculateHash()
	}
	fmt.Println("Block mined:", b.Hash)
}

func (bc *Blockchain) ValidateChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			fmt.Println("Current hashes not equal")
			return false
		}

		if currentBlock.PrevHash != prevBlock.Hash {
			fmt.Println("Previous hashes not equal")
			return false
		}
	}
	return true
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("First Block after Genesis")
	bc.AddBlock("Second Block after Genesis")
	bc.AddBlock("Third Block after Genesis")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %s\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}

	fmt.Println("Blockchain valid:", bc.ValidateChain())
}
