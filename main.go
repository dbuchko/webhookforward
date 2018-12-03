package main
 
import (
    "os"
    "io/ioutil"
    "bytes"
    "fmt"
    "log"
    "net/http"
//    "net/url"
   
)

const (
      DEFAULT_PORT = "8080"
)
type test_struct struct {
     Test string
} 

func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    } 
    switch r.Method {
        case "POST":

 	fmt.Printf("Received a post. forwarding \n");
        forwardurl := os.Getenv("FORWARD_URL")
        fmt.Println("URL read from FORWARD_URL:>", forwardurl)

        reqBody, err := ioutil.ReadAll(r.Body)
           if err !=nil {panic(err)}  
        fmt.Println("The body was:",string(reqBody))
        var  postJson = "{\"text\":\""+string(reqBody)+"\"}"   
//	var postJson = payload={"text": "This is a line of text in a channel.\nAnd this is another line of text."}
        fmt.Println("Sending json: %s\n", postJson)
       postContent := bytes.NewBuffer([]byte(postJson))  

    //    data := url.Values{}
    //    data.Set("payload", postJson)
    //    payloadbody := bytes.NewBufferString(data.Encode())

	resp, err := http.Post(forwardurl, "application/json", postContent)
            if err != nil { panic(err) }

        fmt.Printf("Status: %s\n", resp.Status)
        buf, _ := ioutil.ReadAll(resp.Body)
        fmt.Println(string(buf))  
      
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

    http.HandleFunc("/", hello)
 

    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}