package main

import (
    "flag"
    "fmt"
//    "image/gif"
    "io"
	"net/http"
	"log"
)

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

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    fmt.Printf("%+v\n", w)
    fmt.Printf("%+v\n", req)
	io.WriteString(w, "hello, world!\n")
}

func AvailableImages(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Nothing available yet!\n")
}

func main() {
    SetupConfig()
    
	http.HandleFunc("/hello", HelloServer)
	//http.HandleFunc("/", AvailableImages)
    fmt.Printf("%+v\n", config)
	log.Printf("About to listen on http://127.0.0.1:" + config.port + "/")
	err := http.ListenAndServe(":" + config.port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
