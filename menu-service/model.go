package main

type Restaurant struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
	City    string `json:"city" db:"city"`
}

type MenuItem struct {
	ID       string  `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	Category string  `json:"category" db:"category"`
	Cuisine  string  `json:"cuisine" db:"cuisine"`
	Price    float64 `json:"price" db:"price"`
	RestoID  string  `json:"restaurant_id" db:"id"`
}

type IDResponse struct {
	ID string `json:"id"`
}
