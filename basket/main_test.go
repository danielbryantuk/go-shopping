package main

import (
	"testing"
	"reflect"
)

func TestUpdateBasket(t *testing.T) {
	updateBasket("1", "2", 3)
	expected := Basket{"1", map[string]int{"2":3}}
	actual := basketStore["1"]

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != v%", actual, expected)
	}
}
