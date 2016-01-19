package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Basket struct {
	UserId   string `json:"userId"`
	products map[string]int32 `json:products`
}

//todo - is int32 appropriate here?
var basketStore = make(map[string]map[string]int32)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", healthHandler)
	r.HandleFunc("basket/{userId}", basketHandler)
	log.Fatal(http.ListenAndServe(":3020", r))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func basketHandler(w http.ResponseWriter, r *http.Request) {
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
