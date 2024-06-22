package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

// NewBlock creates and returns Block instances
func NewBlock(data string, prevBlock Block) Block {
	t := time.Now()
	block := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: t.String(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	block.Hash = block.calculateHash()
	return block
}

// calculateHash returns a SHA256 hash of a Block
func (b *Block) calculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + strconv.Itoa(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

// isValidHash checks if the block's hash meets the PoW requirement
func (b *Block) isValidHash(difficulty int) bool {
	prefix := fmt.Sprintf("%0*s", difficulty, "")
	return b.Hash[:difficulty] == prefix
}

// mine performs the proof-of-work computation
func (b *Block) mine(difficulty int) {
	for !b.isValidHash(difficulty) {
		b.Nonce++
		b.Hash = b.calculateHash()
	}
}
