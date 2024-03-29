package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
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
