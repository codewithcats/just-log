package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var listenAddr = flag.String("addr", ":8082", "Address the server will listen on")

func init() {
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(_ http.ResponseWriter, req *http.Request) {
		log.Printf(req.URL.Path)
	})

	fmt.Printf("I am going to listen on %v\n", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
