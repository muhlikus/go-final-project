package main

import (
	"fmt"
	"net/http"
	"os"
)

const webDir = "./web/"

var serverPort string = "7540"

func main() {

	// check if "TODO_PORT" set
	if port, ok := os.LookupEnv("TODO_PORT"); ok {
		//log.Printf("set listening port from ENV to: %s", port)
		serverPort = port
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(webDir)))

	err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), mux)
	if err != nil {
		panic(err)
	}
}
