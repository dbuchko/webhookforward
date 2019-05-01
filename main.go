package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
)

const (
      DEFAULT_PORT = "8080"
)

func handleForward(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
    switch r.Method {
        case "POST":

 	      fmt.Printf("Received a post. forwarding \n");
        forwardurl := os.Getenv("FORWARD_URL")
        fmt.Println("URL read from FORWARD_URL:>", forwardurl)

	      resp, err := http.Post(forwardurl, "application/json", r.Body)
            if err != nil { panic(err) }

        fmt.Printf("Status: %s\n", resp.Status)

    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}

func main() {

    var port string
    if port = os.Getenv("PORT"); len(port) == 0 {
       log.Printf("Warning, PORT not set. Defaulting to %+vn", DEFAULT_PORT)
        port = DEFAULT_PORT
    }

    http.HandleFunc("/", handleForward)


    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
