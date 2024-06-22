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

	bc := pow.NewBlockchain()
	bc.AddBlock("First Block after Genesis", difficulty)
	bc.AddBlock("Second Block after First", difficulty)
}
