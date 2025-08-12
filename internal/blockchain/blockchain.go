/* Core blockchain logic
	Genesis Block Creation (only once)
	Block Creation
	Verification 
	Storage
*/ 

/*
GenesisBlock() -> function to fill out a genesis block
NewBlockchain() -> creates an empty blockchain with a genesis block: returns a blockchain pointer
GetLatestBlock() -> returns the latest block in the blockchain
AddCertification() -> appends a new block containing certificate data: returns the newly added block  
IsValidChain() -> checks the chain integrity: returns true / false
*/

/*
Data Sending:
	internal/storage: for saving blockchain state
	internal/api: for responding to api calls
*/

package main 

import (
	"time"
	"errors"
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

func NewBlockchain() (*Blockchain) {
	var bc Blockchain

	GenesisBlock := CreateGenesisBlock()
	bc.Blocks = append(bc.Blocks, GenesisBlock)
	
	return &bc
}

func GetLatestBlock(bc *Blockchain) Block {
	return bc.Blocks[len(bc.Blocks)-1]	
}

func AddCertification(bc *Blockchain, cert Certificate) (Block, error) {
	latest := GetLatestBlock(bc)
	nb := NewBlock(latest, cert)

	if IsBlockValid(nb, latest) {
		bc.Blocks = append(bc.Blocks, nb)

		// save after adding 
		if err := SaveBlockchain(); err != nil {
			return nb, err
		}
	} 
	return nb, nil
}

func IsValidChain(bc *Blockchain) bool {
	for i := 1; i < len(bc.Blocks); i++ {
		if !IsBlockValid(bc.Blocks[i], bc.Blocks[i-1]) {return false}
	}
	return true
}