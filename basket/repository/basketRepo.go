package repository

type Basket struct {
	UserId   string `json:"userId"`
	Products map[string]int `json:"products"`
}

type BasketRepo interface {
	UpdateBasket(userId string, productId string, quantity int)
	GetStoreAsMap() (map[string]Basket)
	GetBasket(userId string) (Basket, bool)
	SetBasket(userId string, basket Basket)
}
