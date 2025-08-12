/*
	Defines our block structure 
*/

/*
	NewBlock() -> returns a new block object with a valid hash
                  sends data to blockchain.go to append to the blockchain 
	CalculateHash() -> generates a SHA-256 hash of the block's data
                       sends data to NewBlock() and isBlockValid() in blockchain.go			\
	isBlockValid() -> returns a boolean if block is valid in sequence and hash  
					  sends data to blockchain.go for chain validation		
*/

package main 

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Certificate struct {
	MemberID string
	Name string  
	EventName string  
	DateIssued time.Time
}

type Block struct {
	Index int // block's position
	Timestamp time.Time // when block was created
	Data Certificate // certification records
	PrevHash string // hash of previous block
	Hash string  / hash of current block 
}

/* serialize the block */
/* calculate SHA256 hash of the block and convert it to hexadecimal string */
func CalculateHash(block Block) string {
	hash := sha256.Sum256(block)
	hashString := hex.EncodeToString(hash[:])

	return hashString
}

/* fill out newblock's info*/
func NewBlock(previous Block, cert Certificate) Block {
	var newBlock Block

	newBlock.Index = previous.Index + 1
	newBlock.Timestamp = time.Now()
	newBlock.Data = cert 
	newBlock.PrevHash = previous.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}