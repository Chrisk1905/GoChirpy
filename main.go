package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	fmt.Println("creating http server...")

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("serving on port: %s", port)
	log.Fatal(server.ListenAndServe())

}
