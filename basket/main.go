package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"os"
)

var basketService = "localhost:" + os.Getenv("BASKET_SERVICE_PORT")

type Basket struct {
	UserId   string `json:"userId"`
	products map[string]int32 `json:"products"`
}

//userId, Basket
var basketStore = make(map[string]Basket)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", healthHandler)
	r.HandleFunc("/baskets/{userId}", viewBasketHandler).Methods("GET")
	r.HandleFunc("/baskets/{userId}/add", addToBasketHandler).Methods("GET")
	r.HandleFunc("/baskets", updateBasketHandler).Methods("POST")
	r.HandleFunc("/baskets", viewAllBasketsHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(basketService, r))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func viewBasketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("viewBasketHandler entry")
	vars := mux.Vars(r)
	userId := vars["userId"]
	basket, ok := basketStore[userId]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		b, err := json.Marshal(basket)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(b)
	}
}

func updateBasketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("updateBasketHandler entry")
	decoder := json.NewDecoder(r.Body)
	basket := Basket{}
	err := decoder.Decode(&basket)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		basketStore[basket.UserId] = basket
	}
}

func viewAllBasketsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("viewAllBasketsHandler entry")
	b, err := json.Marshal(basketStore)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(b)
	}
}

//todo - very much WIP!
func addToBasketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("addToBasketHandler entry")
	vars := mux.Vars(r)
	userId := vars["userId"]

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		productId := r.FormValue("productId")
		quantity := r.FormValue("quantity")
		fmt.Printf("%v - %v : %v", userId, productId, quantity)
		if val, ok := basketStore[userId]; ok {
			//we have a basket
			//update content
			if basket,ok := val[productId]; ok {
				//todo
			}
		} else {
			//create products map
			var products map[string]int32
			//create new basket
			basketStore[userId] = products
}
}
}

