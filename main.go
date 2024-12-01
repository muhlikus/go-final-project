package main

import "net/http"

const (
	serverPort = ":7540"
	webDir     = "./web/"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(webDir)))

	err := http.ListenAndServe(serverPort, mux)
	if err != nil {
		panic(err)
	}
}
