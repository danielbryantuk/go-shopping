package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	repo "github.com/daniel-bryant-uk/go-shopping/basket/repository"
)

var basketService = "localhost:" + os.Getenv("BASKET_SERVICE_PORT")

var repository repo.LocalBasketStore

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
	basket, ok := repository.GetBasket(userId)
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
	basket := repo.Basket{}
	err := decoder.Decode(&basket)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		repository.SetBasket(basket.UserId, basket)
	}
}

func viewAllBasketsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("viewAllBasketsHandler entry")
	log.Printf("%v\n", repository.GetStoreAsMap())
	b, err := json.Marshal(repository.GetStoreAsMap())
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
		quantityStr := r.FormValue("quantity")
		quantity, _ := strconv.Atoi(quantityStr)

		fmt.Printf("UserId: %v - Adding product %v, with quantity %v\n", userId, productId, quantity)
		repository.UpdateBasket(userId, productId, quantity)

		w.WriteHeader(http.StatusOK)
	}
}



