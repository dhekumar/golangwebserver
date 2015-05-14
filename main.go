package main

import (
	"golangwebserver/constants"
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	// serve static assets showing how to strip/change the path.
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("resources/images"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("resources/javascripts"))))
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("resources/stylesheets"))))

	log.Fatal(http.ListenAndServe(constants.HostName, router))
	log.Printf("Listening - %s\n", constants.HostName)

}
