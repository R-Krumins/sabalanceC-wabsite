package routers

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Name  string  `json:"name"`
	Desc  string  `json:"desc"`
	Price float32 `json:"price"`
	Img   string  `json:"img"`
}

func ProductRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/product", getProductsHandler)
	return mux
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product = getProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getProducts() []Product {
	return []Product{
		{
			Name:  "Organic Quinoa",
			Desc:  "Premium organic quinoa, rich in protein and essential nutrients",
			Price: 8.99,
			Img:   "https://www.nutritionadvance.com/wp-content/uploads/2018/01/several-fresh-carrots-with-intact-green-stems.jpg",
		},
		{
			Name:  "Organic Avocados",
			Desc:  "Farm-fresh organic avocados, perfectly ripe and ready to eat",
			Price: 6.99,
			Img:   "https://www.nutritionadvance.com/wp-content/uploads/2018/01/several-fresh-carrots-with-intact-green-stems.jpg",
		},
		{
			Name:  "Organic Raw Honey",
			Desc:  "Pure unfiltered organic honey harvested from sustainable apiaries",
			Price: 12.99,
			Img:   "https://www.nutritionadvance.com/wp-content/uploads/2018/01/several-fresh-carrots-with-intact-green-stems.jpg",
		},
		{
			Name:  "Organic Kale",
			Desc:  "Fresh organic kale, locally grown without pesticides",
			Price: 3.99,
			Img:   "https://www.nutritionadvance.com/wp-content/uploads/2018/01/several-fresh-carrots-with-intact-green-stems.jpg",
		},
		{
			Name:  "Organic Chia Seeds",
			Desc:  "Nutrient-dense organic chia seeds, perfect for smoothies and baking",
			Price: 8.49,
			Img:   "https://www.nutritionadvance.com/wp-content/uploads/2018/01/several-fresh-carrots-with-intact-green-stems.jpg",
		},
	}
}
