package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
)

var sampleProduct Product

func init() {
	sampleProduct = Product{"1", "12345", "Widget", "Premium brand widgets", 120}
}

type Product struct {
	Id          string
	Sku         string
	Name        string
	Description string
	Price       int64
}

func main() {
	//TODO - Do we need Gorilla for this example??
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/product", productHandler)
	log.Fatal(http.ListenAndServe(":3000", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", sampleProduct)
}