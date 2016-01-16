package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

var products map[int32]Product

func init() {
	products = createProducts()
}

type Product struct {
	Id          string `json:"id"`
	Sku         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64 `json:"price"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products/{productId}", handleAllProducts)
	router.HandleFunc("/products/{productId}", handleProduct)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func handleAllProducts(w http.ResponseWriter, r *http.Request) {

}

func handleProduct(w http.ResponseWriter, r *http.Request) {

}

func createProducts() map[int32]Product {
	return map[int32]Product{
		1: Product{"1", "12345", "Widget", "Premium Widget", 120},
	}
}

