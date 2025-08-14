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
    "fmt"
    "html/template"
    "net/http"
    "time"
    "github.com/jackc/pgx/v5"
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
)

// package level variables
var blockchainPtr *Blockchain  
var dbPool *pgx.Pool

func ConnectDB(databaseURL string){
    var err error 
    dbPool, err = pgxpool.New(context.Background(), databaseURL)
    if err != nil {
        log.Fatalf("trouble connecting to database: %v\n", err)
    }
    fmt.Println("connected to postgresql database")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
        return
    }
    tmpl, err := template.ParseFiles("../web/form.html") 
    if err != nil {
        http.Error(w, "unable to load form", http.StatusInternalServerError)
        return
    }
    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, "unable to render form", http.StatusInternalServerError)
        return
    }
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
        return
    }
    if err := r.ParseForm(); err != nil {
        http.Error(w, "unable to parse form", http.StatusBadRequest)
        return
    }
    attendeeName := r.FormValue("name")
    eventName := r.FormValue("event")

    if attendeeName == "" || eventName == "" {
        http.Error(w, "missing required fields", http.StatusBadRequest)
        return
    }

    cert := Certificate{
        MemberID:  "", // generate or assign if needed
        Name:      attendeeName,
        EventName: eventName,
        DateIssued: time.Now(),
    }

    _, err := AddCertification(blockchainPtr, cert)
    if err != nil {
        http.Error(w, "unable to add certificate", http.StatusInternalServerError)
        return
    }

    if err := SaveBlockchain("blockchain.json", blockchainPtr); err != nil {
        http.Error(w, "unable to save blockchain", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, "<h1>Thank you, %s!</h1><p>Your attendance at '%s' has been recorded.</p>", attendeeName, eventName)
}

func main() {

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        <-stop
        fmt.Println("shutting down server...")
        dbPool.Close()
        os.Exit(0)
    }()

    var err error
    blockchainPtr, err = LoadBlockchain("blockchain.json")
    if err != nil {
        panic(err)
    }

    // Add root route handler that redirects to /form
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            http.Redirect(w, r, "/form", http.StatusSeeOther)
        } else {
            http.NotFound(w, r)
        }
    })

    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/submit", submitHandler)

    fmt.Println("Server starting at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
    fmt.Printf("working")
}

