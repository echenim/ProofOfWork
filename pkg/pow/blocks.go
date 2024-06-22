package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PrevHash     string
	Hash         string
	Nonce        int
	Transactions []Transaction
}

// NewBlock creates and returns Block instances
func NewBlock(data string, prevBlock Block, transactions []Transaction) Block {
	t := time.Now()
	block := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: t.String(),
		Data:      data,
		Transactions: transactions,
		PrevHash:  prevBlock.Hash,
	}
	block.Hash = block.calculateHash()
	return block
}

// calculateHash returns a SHA256 hash of a Block
func (b *Block) calculateHash() string {
	records := strconv.Itoa(b.Index) + b.Timestamp + b.PrevHash + strconv.Itoa(b.Nonce)
	for _, tx := range b.Transactions {
		records += tx.Sender + tx.Receiver + fmt.Sprintf("%f", tx.Amount) // Simple concatenation of transaction details
	}
	h := sha256.New()
	h.Write([]byte(records))
	return hex.EncodeToString(h.Sum(nil))
}

// isValidHash checks if the block's hash meets the PoW requirement
func (b *Block) isValidHash(difficulty int) bool {
	prefix := fmt.Sprintf("%0*s", difficulty, "")
	return b.Hash[:difficulty] == prefix
}

// mine performs the proof-of-work computation
func (b *Block) mine(difficulty int) {
	if !b.areTransactionsValid() {
		log.Println("Block contains invalid transactions")
		return
	}
	for {
		b.Nonce++
		b.Hash = b.calculateHash()
		if b.isValidHash(difficulty) {
			break
		}
	}
}

func (b *Block) areTransactionsValid() bool {
	for _, tx := range b.Transactions {
		if !tx.IsValid() {
			return false
		}
	}
	return true
}
