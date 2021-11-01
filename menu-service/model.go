package main

type Restaurant struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
	City    string `json:"city" db:"city"`
}

type IDResponse struct {
	ID string `json:"id"`
}
