/*
save and load blockchain data to disk so it persists between runs 
*/

/*
SaveBlockchain() -> saves blockchain to json file; returns error if any 
LoadBlockchain() -> reads blockchain from json file; returns blockchain object

sends data to:
	interanl/blockchain.go: when loading
	called by: internal/blockchain after adding blocks
*/

package main 

import (
	"encoding/json"
	"os"
	"errors"
	"fmt"
)

func SaveBlockChain(filename string, bc *Blockchain) error {
	// marshal bc to json
	JSONbc, err := json.Marshal(bc)
	if err != nil {return err}

	// write json bytes to the file
	err = os.WriteFile(filename, JSONbc, 0644) 
	if err != nil {return err}

	return nil // success 
}

func LoadBlockchain(filename string) (bc *Blockchain, error) {
	bc_file, err := os.ReadFile(filename)
	
	// logic for file not found (blockchain doesn't exist)
	if err != nil {return err}

	_, err = os.Stat(filename)
	if err != nil {
		NewBC := NewBlockchain()
		return NewBC
	} else {
		err := json.Unmarshal(filename, &bc_file)
		return bc
	} 
	fmt.Printf(">> IN storage.go << : error reading or parsing data")
	return err
}