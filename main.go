package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(_ http.ResponseWriter, req *http.Request) {
		log.Printf(req.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}
