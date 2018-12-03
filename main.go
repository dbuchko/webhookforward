package main
 
import (
    "os"
    "io/ioutil"
    "bytes"
    "fmt"
    "log"
    "net/http"
    "encoding/json"

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

        reqBody, err := ioutil.ReadAll(r.Body)
           if err !=nil {panic(err)}  
        fmt.Println("The body was:",string(reqBody))
        
   
         var f interface{}
         err3 := json.Unmarshal(reqBody,&f)
                 if err3 !=nil { panic (err) }
         m := f.(map[string]interface{})

         messageSubject := m["subject"]
       //messageBody := m["body"]
         messageTopic := m["topic"]
         messageTimestamp := m["timestamp"]

         var  postJson = "{\"text\":\" Subject:"+messageSubject.(string)+"\n Time:"+messageTimestamp.(string)+ " \n Topic:"+messageTopic.(string)+"\"}"   
        fmt.Println("Sending json: %s\n", postJson)
        postContent := bytes.NewBuffer([]byte(postJson))  
   
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

    http.HandleFunc("/", handleForward)
 

    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
