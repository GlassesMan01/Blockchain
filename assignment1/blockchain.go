package blockchain

import (
    "crypto/sha256"
    "fmt"
    "strings"
)

// Block structure
type Block struct {
    Transaction  string
    Nonce        int
    PreviousHash string
    CurrentHash  string
}

// BlockList structure
type BlockList struct {
    List []*Block
}

// NewBlock creates a new block
func NewBlock(transaction string, nonce int, previousHash string) *Block {
    b := new(Block)
    b.Transaction = transaction
    b.Nonce = nonce
    b.PreviousHash = previousHash
    b.CurrentHash = CalculateHash(transaction, nonce, previousHash)
    return b
}

// AppendBlock adds a new block to the blockchain
func (bl *BlockList) AppendBlock(transaction string, nonce int) *Block {
    var previousHash string
    if len(bl.List) > 0 {
        previousHash = bl.List[len(bl.List)-1].CurrentHash
    }
    newBlock := NewBlock(transaction, nonce, previousHash)
    bl.List = append(bl.List, newBlock)
    return newBlock
}

// ListBlocks displays all the blocks in the blockchain
func (bl *BlockList) ListBlocks() {
    for i, block := range bl.List {
        fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
        fmt.Printf("Transaction: %s \n", block.Transaction)
        fmt.Printf("Nonce: %d \n", block.Nonce)
        fmt.Printf("Previous Block Hash: %s \n", block.PreviousHash)
        fmt.Printf("Current Block Hash: %s \n", block.CurrentHash)
        fmt.Println(strings.Repeat("=", 60))
    }
}

// ChangeBlock updates the transaction of a specific block
func (bl *BlockList) ChangeBlock(id int, newTransaction string) {
    block := bl.List[id]
    block.Transaction = newTransaction
    block.CurrentHash = CalculateHash(newTransaction, block.Nonce, block.PreviousHash)
    fmt.Printf("Block %d transaction changed to: %s \n", id, newTransaction)
}

// VerifyChain verifies the integrity of the blockchain
func (bl *BlockList) VerifyChain() {
    for i := 1; i < len(bl.List); i++ {
        previousBlock := bl.List[i-1]
        currentBlock := bl.List[i]
        expectedPreviousHash := CalculateHash(previousBlock.Transaction, previousBlock.Nonce, previousBlock.PreviousHash)
        if currentBlock.PreviousHash != expectedPreviousHash {
            fmt.Printf("Blockchain is invalid at Block: %d \n", i)
            return
        }
    }
    fmt.Println("Blockchain is valid! \n")
}

// CalculateHash computes the hash of a block's data
func CalculateHash(transaction string, nonce int, previousHash string) string {
    data := fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)
    hash := sha256.Sum256([]byte(data))
    return fmt.Sprintf("%x", hash)
}
