package main

import (
    "flag"
    "fmt"
//    "image/gif"
    "io"
	"net/http"
	"log"
)

// Start of response structures //
type Error struct {
    code int
    message string
}

type ResponseBody struct {
}

type Response struct {
    statusCode int
    error Error
    body string
}

func (resp *Response) setError(errorCode int, message string) {
    error := Error{code: errorCode, message: message}
    resp.error = error
    resp.statusCode = http.StatusBadRequest
}

func (resp *Response) setResponse(body string) {
    resp.body = body
    resp.statusCode = 200
}

func (resp *Response) send(w http.ResponseWriter) {
    if resp.error.code != 0 {
        http.Error(w, resp.error.message, resp.statusCode)
    }
    
	io.WriteString(w, resp.body)
}
// End of response structures //

// Start of app config //
type Config struct {
    port string
}

var config Config;

// Load and parse all the available flags
func SetupConfig() {
    var port string
    flag.StringVar(&port, "port", "12345", "-port=\"12345\" to set a port to listen on")
    flag.Parse()
    
    config = Config{port: port}
}
// End of app config //

// Simple ping for health check
func HandlePing(w http.ResponseWriter, req *http.Request) {
    resp := new(Response)
    resp.setResponse("Ping!")
    resp.send(w)
}

func HandleImg(w http.ResponseWriter, req *http.Request) {
	resp := new(Response)
    resp.setError(1000, "Nothing available yet!")
    resp.send(w)
}

func HandleRoot(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Nothing available yet!\n")
}

func main() {
    SetupConfig()
    
	http.HandleFunc("/ping", HandlePing)
	http.HandleFunc("/img", HandleImg)
	http.HandleFunc("/", HandleRoot)
    
    // Set up the port based on the flag passed in when invoking this app
    fmt.Printf("%+v\n", config)
	log.Printf("About to listen on http://127.0.0.1:" + config.port + "/")
	err := http.ListenAndServe(":" + config.port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
