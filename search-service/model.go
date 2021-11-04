package main

import "encoding/json"

type MenuItem struct {
	ID       string  `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	Category string  `json:"category" db:"category"`
	Cuisine  string  `json:"cuisine" db:"cuisine"`
	Price    float64 `json:"price" db:"price"`
	RestoID  string  `json:"restaurant_id" db:"id"`
}

// MarshalJSON called during insert into elasticsearch
func (mi MenuItem) MarshalJSON() ([]byte, error) {
	fields := map[string]interface{}{
		"name":          mi.Name,
		"category":      mi.Category,
		"cuisine":       mi.Cuisine,
		"price":         mi.Price,
		"restaurant_id": mi.RestoID,
	}
	return json.Marshal(fields)
}
