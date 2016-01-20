package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

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
	r.HandleFunc("/baskets", updateBasketHandler).Methods("POST")
	r.HandleFunc("/baskets", viewAllBasketsHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":3020", r))
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
