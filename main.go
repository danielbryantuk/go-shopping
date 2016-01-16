package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func main() {
	//TODO - Do we need Gorilla for this example??
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":3000", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}