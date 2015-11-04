// +build OMIT

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHello) // HL
	fmt.Println("serving ResponseWriter http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe("localhost:8080", nil)) // HL
}

func handleHello(w http.ResponseWriter, req *http.Request) { // HL
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "Hello Codecamp!")
}
