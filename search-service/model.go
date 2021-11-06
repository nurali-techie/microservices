package main

type Restaurant struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	City    string `json:"city,omitempty"`
}

type MenuItem struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Cuisine  string  `json:"cuisine"`
	Price    float64 `json:"price"`
	RestoID  string  `json:"restaurant_id"`
}

type SearchResult struct {
	MenuItem
	RestoName string `json:"restaurant_name"`
}

type MenuItemES struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Cuisine  string  `json:"cuisine"`
	Price    float64 `json:"price"`
	RestoID  string  `json:"restaurant_id"`
}

func (m *MenuItem) ToMenuItemES() *MenuItemES {
	return &MenuItemES{
		Name:     m.Name,
		Category: m.Category,
		Cuisine:  m.Cuisine,
		Price:    m.Price,
		RestoID:  m.RestoID,
	}
}
