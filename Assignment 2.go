package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Structure of Block
type Block struct {
	Timestamp    string
	Nonce        int
	PreviousHash string
	CurrentHash  string
	Transactions string
}

// Structure of Transaction
type Transaction struct {
	TransactionId              string  `json:"transaction_id"`
	SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
	RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
	Value                      float32 `json:"value"`
}

// List of Blocks
type Blockchain struct {
	Chain           []*Block
	TransactionPool []*Transaction
}

// Calculate Hash of Block
func calculateHash(stringValue1 string, numericValue int, stringValue2 string) string {
	data := fmt.Sprintf("%s%d%s", stringValue1, numericValue, stringValue2)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// Add new block
func newBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	now := time.Now().String()
	b.Timestamp = now
	b.Nonce = nonce
	b.PreviousHash = previousHash
	b.CurrentHash = calculateHash(b.Timestamp, nonce, previousHash)
	return b
}

// Append the block in the List
func (bl *Blockchain) appendBlock(difficulty int) *Block {
	//for previous hash of block
	var previousHash string
	if len(bl.Chain) > 0 {
		previousHash = bl.Chain[len(bl.Chain)-1].CurrentHash
	} else {
		previousHash = "0000"
	}

	nonce := proofOfWork(previousHash, difficulty)

	//adding new block
	addBlock := newBlock(nonce, previousHash)

	//Add transactions in the Block in JSON format
	transactionsData, err := json.Marshal(bl.TransactionPool)
	if err != nil {
		fmt.Println("Error in adding marshall transactions: ", err)
	}

	//adding transactions in the block
	addBlock.Transactions = string(transactionsData)

	//clear the transaction pool
	bl.TransactionPool = nil

	//append the block
	bl.Chain = append(bl.Chain, addBlock)
	return addBlock
}

// List all Blocks
func (bl *Blockchain) listBlocks() {
	for i, block := range bl.Chain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 40), i, strings.Repeat("=", 40))
		fmt.Printf("Timestamp: %s \n", block.Timestamp)
		fmt.Printf("Nonce: %d \n", block.Nonce)
		fmt.Printf("Previous Block Hash: %s \n", block.PreviousHash)
		fmt.Printf("Current Block Hash: %s \n", block.CurrentHash)
		fmt.Println(strings.Repeat("-", 89))
		fmt.Printf("Transactions: \n%s\n", block.Transactions)
		fmt.Println(strings.Repeat("=", 89))
		fmt.Println("\n")
	}
}

// Print a specific Block by index
func (bl *Blockchain) printBlock(index int) {
	// Check if the index is within bounds
	if index < 0 || index >= len(bl.Chain) {
		fmt.Println("Error: Block index out of range.")
		return
	}

	// Retrieve the block at the specified index
	block := bl.Chain[index]

	// Print the block's details
	fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 40), index, strings.Repeat("=", 40))
	fmt.Printf("Timestamp: %s \n", block.Timestamp)
	fmt.Printf("Nonce: %d \n", block.Nonce)
	fmt.Printf("Previous Block Hash: %s \n", block.PreviousHash)
	fmt.Printf("Current Block Hash: %s \n", block.CurrentHash)
	fmt.Println(strings.Repeat("-", 89))
	fmt.Printf("Transactions: \n%s\n", block.Transactions)
	fmt.Println(strings.Repeat("=", 89))
	fmt.Println("\n")
}

// New transaction
func newTransaction(sender string, recipient string, value float32) *Transaction {
	tr := new(Transaction)
	tr.SenderBlockchainAddress = sender
	tr.RecipientBlockchainAddress = recipient
	tr.Value = value
	tr.TransactionId = calculateHash(sender, int(value), recipient)
	return tr
}

// append transaction
func (bl *Blockchain) appendTransaction(sender string, recipient string, value float32) *Transaction {
	addTransaction := newTransaction(sender, recipient, value)
	bl.TransactionPool = append(bl.TransactionPool, addTransaction)
	return addTransaction
}

// Proof of Work: find nonce
func proofOfWork(previousHash string, difficulty int) int {
	nonce := 0
	prefix := strings.Repeat("0", difficulty)
	for {
		hash := calculateHash(time.Now().String(), nonce, previousHash)
		if strings.HasPrefix(hash, prefix) {
			break
		}
		nonce++
	}
	return nonce
}

func main() {
	blockchain := new(Blockchain)

	//add some transactions
	blockchain.appendTransaction("Alice", "Bob", 10.5)
	blockchain.appendTransaction("Bob", "Charlie", 20.0)

	//append block with proof of work difficulty level 2
	blockchain.appendBlock(2)

	//add new transactions for new block
	blockchain.appendTransaction("Charlie", "John", 10.5)
	blockchain.appendTransaction("John", "Alice", 20.0)

	//append block with proof of work difficulty level 2
	blockchain.appendBlock(2)

	//blockchain.listBlocks()

	blockchain.printBlock(1)
}
