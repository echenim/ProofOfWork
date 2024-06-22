package main

import (
	"log"
	"os"
	"strconv"

	"github.com/echenim/ProofOfWork/pkg/pow"
	"github.com/joho/godotenv"
)

func setupLogging() {
	logFile, err := os.OpenFile("blockchain.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	log.SetOutput(logFile)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	setupLogging()
	difficulty := 4
	if diff, exists := os.LookupEnv("DIFFICULTY"); exists {
		var err error
		difficulty, err = strconv.Atoi(diff)
		if err != nil {
			log.Printf("Invalid difficulty level provided: %v. Using default.\n", diff)
		}
	}

	// Initialize blockchain
	bc := pow.NewBlockchain()

	// Define transactions for the first block
	txs1 := []pow.Transaction{
		{Sender: "Alice", Receiver: "Bob", Amount: 10},
		{Sender: "Charlie", Receiver: "Dave", Amount: 15},
	}

	// Add first block with transactions
	bc.AddBlock("First Block after Genesis", txs1, difficulty)

	// Define transactions for the second block
	txs2 := []pow.Transaction{
		{Sender: "Eve", Receiver: "Frank", Amount: 5},
		{Sender: "Gina", Receiver: "Harry", Amount: 20},
	}

	// Add second block with transactions
	bc.AddBlock("Second Block after First", txs2, difficulty)
}
