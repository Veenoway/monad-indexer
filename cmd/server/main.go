package main 

import (
    "fmt"
)

type Transaction struct {
	Hash string `json:"hash"`
	From string `json:"from"`
	To string `json:"to"`
	Value uint64 `json:"value"`
}

type Block struct {
	Number uint64 `json:"number"`
	Hash string `json:"hash"`
	Timestamp uint64 `json:"timestamp"`
}


func main() {

	block := Block{
		Number: 1,
		Hash: "9x",
		Timestamp: 1717000000,
	}

	transaction := Transaction{
		Hash: "9x",
		From: "0x1234567890",
		To: "0x1234567890",
		Value: 100,
	}

	fmt.Println("--------------------------------")
	fmt.Printf("Block: %v\nHash: %v\nTimestamp: %v\n", block.Number, block.Hash, block.Timestamp)
	fmt.Println("--------------------------------")
	fmt.Printf("Hash: %v\nFrom: %v\nTo: %v\nValue: %v\n", transaction.Hash, transaction.From, transaction.To, transaction.Value)
	fmt.Println("--------------------------------")

}
