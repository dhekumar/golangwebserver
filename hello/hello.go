package main

import (
	"fmt"
	"log"
	"net/http"
	"string"
)

type WebServer struct{}

func (web WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var w WebServer
	fmt.Println(string.Reverse("hello world program"))
	err := http.ListenAndServe("localhost:4000", w)

	if err != nil {
		log.Fatal(err)
	}
}
