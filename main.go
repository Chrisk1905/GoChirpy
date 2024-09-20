package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("creating http server...")

	var serveMux *http.ServeMux = http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

}
