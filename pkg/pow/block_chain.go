package pow

import (
	"log"
	"sync"
	"time"
)

// Blockchain is a series of validated Blocks
type Blockchain struct {
	blocks []Block
	mutex  sync.Mutex
}

// AddBlock handles adding a block to the blockchain
func (bc *Blockchain) AddBlock(data string, transactions []Transaction, difficulty int) {
	if len(transactions) == 0 {
		log.Println("No transactions to add")
		return
	}

	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	lastBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, lastBlock, transactions)
	if !newBlock.areTransactionsValid() {
		log.Println("Attempt to add invalid transactions to the block")
		return
	}
	newBlock.mine(difficulty)
	bc.blocks = append(bc.blocks, newBlock)
	log.Printf("Block %d added to blockchain\n", newBlock.Index)
	log.Printf("Hash: %s\n", newBlock.Hash)
}

func NewBlockchain() *Blockchain {
	genesisBlock := Block{
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
	}
	genesisBlock.Hash = genesisBlock.calculateHash()
	return &Blockchain{blocks: []Block{genesisBlock}}
}
