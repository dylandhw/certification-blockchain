/*
	Defines our block structure 
*/

/*
	NewBlock() -> returns a new block object with a valid hash
                  sends data to blockchain.go to append to the blockchain 
	CalculateHash() -> generates a SHA-256 hash of the block's data
                       sends data to NewBlock() and isBlockValid() in blockchain.go			\
	IsBlockValid() -> returns a boolean if block is valid in sequence and hash  
					  sends data to blockchain.go for chain validation		
*/

package main 

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
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
	Timestamp string // when block was created
	Data Certificate // certification records
	PrevHash string // hash of previous block
	Hash string `json:"-"` // hash of current block 
}

/* serialize the block */
/* calculate SHA256 hash of the block and convert it to hexadecimal string */
func CalculateHash(block Block) string {
	blockSerialized, err := json.Marshal(block)
	if err != nil {
		fmt.Printf("error serializing block: %v\n", err)
		return ""
	}

	hash := sha256.Sum256(blockSerialized)
	hashString := hex.EncodeToString(hash[:])

	return hashString
}

/* fill out newblock's info*/
func NewBlock(previous Block, cert Certificate) Block {
	var newBlock Block

	newBlock.Index = previous.Index + 1
	newBlock.Timestamp = time.Now().Format(time.RFC3339)
	newBlock.Data = cert 
	newBlock.PrevHash = previous.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}

/* compares hashes, indexes, and links between the blocks */
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index + 1 != newBlock.Index {return false}
	
	if oldBlock.Hash != newBlock.PrevHash {return false}

	if CalculateHash(newBlock) != newBlock.Hash {return false}

	return true
}