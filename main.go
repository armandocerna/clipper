package main

import (
	"clipper/clipper"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var endpoint, port string

func init() {
	flag.StringVar(&endpoint,"endpoint", "0.0.0.0", "The host to listen to defaults to 0.0.0.0")
	flag.StringVar(&port, "port", "12345", "Port to listen on defaults to 12345")
}

func main() {
	flag.Parse()
	fmt.Printf("starting webserver on %s:%s\n", endpoint, port)
	http.HandleFunc("/", clipper.DashboardHandler)
	go clipper.ReadClipboard()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", endpoint, port), nil))
}
