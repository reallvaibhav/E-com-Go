package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type Promotion struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	DiscountPercentage float64   `json:"discount_percentage"`
	ApplicableProducts []string  `json:"applicable_products"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	IsActive           bool      `json:"is_active"`
}

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var (
	promotions      = make(map[string]Promotion)
	promotionsMutex = sync.RWMutex{}
	products        = []Product{
		{ID: "p1", Name: "T-shirt", Price: 25.00},
		{ID: "p2", Name: "Jeans", Price: 50.00},
		{ID: "p3", Name: "Sneakers", Price: 90.00},
	}
)

func createPromotionHandler(w http.ResponseWriter, r *http.Request) {
	var promo Promotion
	if err := json.NewDecoder(r.Body).Decode(&promo); err != nil {
		http.Error(w, "Invalid promotion data", http.StatusBadRequest)
		return
	}
	promotionsMutex.Lock()
	promotions[promo.ID] = promo
	promotionsMutex.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(promo)
}

func getProductsWithPromotionsHandler(w http.ResponseWriter, r *http.Request) {
	promotionsMutex.RLock()
	defer promotionsMutex.RUnlock()

	productMap := make(map[string]Product)
	for _, p := range products {
		productMap[p.ID] = p
	}

	var promotedProducts []Product
	for _, promo := range promotions {
		if !promo.IsActive {
			continue
		}
		for _, pid := range promo.ApplicableProducts {
			if product, exists := productMap[pid]; exists {
				promotedProducts = append(promotedProducts, product)
			}
		}
	}

	json.NewEncoder(w).Encode(promotedProducts)
}

func deletePromotionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/promotions/"):]

	promotionsMutex.Lock()
	defer promotionsMutex.Unlock()
	if _, exists := promotions[id]; !exists {
		http.Error(w, "Promotion not found", http.StatusNotFound)
		return
	}
	delete(promotions, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/promotions", createPromotionHandler)
	http.HandleFunc("/products-with-promotions", getProductsWithPromotionsHandler)
	http.HandleFunc("/promotions/", deletePromotionHandler)

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
