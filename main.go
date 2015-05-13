package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	// serve static assets showing how to strip/change the path.
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("resources/javascripts"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("resources/javascripts"))))
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("resources/javascripts"))))

	log.Fatal(http.ListenAndServe("localhost:4000", router))

}
