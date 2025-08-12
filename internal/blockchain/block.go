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

type Block struct {
	Index int // block's position
	Timestamp string // when block was created
	Data string // certification records
	PrevHash string // hash of previous block
	Hash string  / hash of current block 
	Nonce int // for proof of work - not sure if this will be used
}