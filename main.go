package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe("localhost:4000", router))
}
