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

func LoadBlockChain(filename string) (*Blockchain, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return NewBlockchain(), nil // start fresh
		}
		return nil, err // some other read error
	}

	var bc Blockchain
	if err := json.Unmarshal(data, &bc); err != nil {
		return nil, err
	}
	fmt.Printf(">>> /storage.go : error reading or parsing json data") // internal debugging
	return &bc, nil
}