package main

import (
	"testing"
	"reflect"
)

func TestUpdateBasket(t *testing.T) {
	//todo - find better way of resetting package-level vars between tests
	basketStore = make(map[string]Basket)

	updateBasket("1", "2", 3)
	expected := Basket{"1", map[string]int{"2":3}}
	actual := basketStore["1"]

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != v%", actual, expected)
	}
}

func TestUpdateBasketTwice(t *testing.T) {
	//todo - find better way of resetting package-level vars between tests
	basketStore = make(map[string]Basket)

	updateBasket("1", "2", 3)
	updateBasket("1", "2", 2)
	expected := Basket{"1", map[string]int{"2":5}}
	actual := basketStore["1"]

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != v%", actual, expected)
	}
}

