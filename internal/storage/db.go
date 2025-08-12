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