/* Core blockchain logic
	Genesis Block Creation (only once)
	Block Creation
	Verification 
	Storage
*/ 

/*
GenesisBlock() -> function to fill out a genesis block
NewBlockchain() -> creates an empty blockchain with a genesis block: returns a blockchain pointer
AddCertification() -> appends a new block containing certificate data: returns the newly added block  
IsValid() -> checks the chain integrity: returns true / false
*/

/*
Data Sending:
	internal/storage: for saving blockchain state
	internal/api: for responding to api calls
*/

package main 

import (
	"time"
)

type Blockchain struct {
	Blocks []Block // slice of block structs 
	// will add concurrent access abilities later on
}

func CreateGenesisBlock() Block {
	var emptyCert Certificate
	var GenesisBlock Block

	GenesisBlock.Index = 0
	GenesisBlock.Timestamp = time.Now().Format(time.RFC3339)
	GenesisBlock.Data = emptyCert
	GenesisBlock.PrevHash = "0"
	GenesisBlock.Hash = CalculateHash(GenesisBlock)

	return GenesisBlock
}

func NewBlockchain() (*Blockchain, error) {
	var Blockchain Blockchain
	BlockchainPtr := &Blockchain

	Genesis := CreateGenesisBlock()

	return BlockchainPtr
}