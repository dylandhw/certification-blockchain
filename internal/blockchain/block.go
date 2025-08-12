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