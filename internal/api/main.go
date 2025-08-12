/* Initializes blockchain, routes, and starts the http server */
/* Passes initialized blockchain object to HTTP handlers */


/*
RegisterRoutes() -> set up routes
handleSubmitCertification() -> accepts qr form submissions, stores pending requests, reeturns a json success message
handleApproveCertification() -> accepts admin approval, adds to blockchain 
handleGetCertifications() -> returns all certifications for a given userID 

sends data to:
	internal/blockchain: to add approved certifications
*/

package main 

import (
	"errors"
)

func main(){
	bc, err := LoadBlockchain("blockchain.json")
	if err != nil {panic(err)}

}