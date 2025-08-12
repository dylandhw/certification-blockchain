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


/*TODO
Create routes for:
	/ -> serves form.html (for adding/searching/ badges)
	/lookup -> serves lookup.html (for viewing a specific badge)
	/submit -> post route to add a new badge to the blockchain
	/qrcode/{id} -> returns a qr code png for the badge
*/

package main 

import (
	"errors"
	"log"
	"net/http"
	"html/template"
)

func formHandler(w http.ResponseWriter, r *http.Request) string {
	if r.method != http.MethodGet {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("./web/form.html")
	if err != nil {
		http.Error(w, " : unable to load form", http.StatusInternalServerError)
		return
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request){}


func main(){
	// load our blockchain
	bc, err := LoadBlockchain("blockchain.json")
	if err != nil {panic(err)}
}
