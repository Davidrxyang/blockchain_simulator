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
