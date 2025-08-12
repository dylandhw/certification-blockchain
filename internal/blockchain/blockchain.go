/* Core blockchain logic
	Block Creation
	Verification 
	Storage
*/ 

/*
NewBlockchain() -> creates an empty blockchain with a genesis block: returns a blockchain pointer
AddCertification() -> appends a new block containing certificate data: returns the newly added block  
IsValid() -> checks the chain integrity: returns true / false
*/

/*
Data Sending:
	internal/storage: for saving blockchain state
	internal/api: for responding to api calls
*/