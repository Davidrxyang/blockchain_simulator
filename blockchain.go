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

