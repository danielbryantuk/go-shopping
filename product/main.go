package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

var products map[string]Product

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
	router.HandleFunc("/products", handleAllProducts)
	router.HandleFunc("/products/{productId}", handleProduct)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func handleAllProducts(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}
	w.Write(b)
}

func handleProduct(w http.ResponseWriter, r *http.Request) {

}

func createProducts() map[string]Product {
	return map[string]Product{
		"1": Product{"1", "12345", "Widget", "Premium Widget", 120},
		"2": Product{"2", "56789", "Splunket", "Basic Splunket", 20},
		"3": Product{"3", "45678", "Bolt", "Premium Bolt", 260},
	}
}

