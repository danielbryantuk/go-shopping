package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"html/template"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

const viewDir = "views/"

//service locations - this would obviously be read from an external source
const storeService = "localhost:3000"
const productService = "http://localhost:3010/"

type Product struct {
	Id          string `json:"id"`
	Sku         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64 `json:"price"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/product", productHandler)
	r.HandleFunc("/view/product", productViewHandler)
	log.Fatal(http.ListenAndServe(storeService, r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	//todo
}

func productViewHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(productService + "products")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var products map[string]Product

	if err := json.Unmarshal(body, &products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(os.Stderr, err)
	} else {
		t, err := template.ParseFiles(viewDir + "shopFront.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(os.Stderr, err)
		} else {
			t.Execute(w, products)
		}
	}
}