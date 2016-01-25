package repository

type Basket struct {
	UserId   string `json:"userId"`
	Products map[string]int `json:"products"`
}

//userId, Basket
var basketStore = make(map[string]Basket)

func UpdateBasket(userId string, productId string, quantity int) {
	//todo return bool indicating success
	if basket, ok := basketStore[userId]; ok {
		//we have a basket for the user
		//update content
		if _, ok := basket.Products[productId]; ok {
			//bump count
			basket.Products[productId] = basket.Products[productId] + quantity
		} else {
			//add one more
			basket.Products[productId] = quantity
			basketStore[userId] = basket
		}

	} else {
		var products = make(map[string]int)
		products[productId] = quantity
		basket = Basket{userId, products}
		basketStore[userId] = basket
	}
}

func GetStoreAsMap() (map[string]Basket) {
	return basketStore
}

func GetBasket(userId string) (Basket, bool) {
	//todo - I don't think this is necessary - just return map access?
	if basket, ok := basketStore[userId]; ok {
		return basket, true
	} else {
		return nil, false
	}
}


func SetBasket(userId string, basket Basket) { //todo - return
	//todo - I don't think this is necessary - just return map access?
	basketStore[userId] = basket
}

