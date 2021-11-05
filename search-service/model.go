package main

type MenuItem struct {
	ID       string  `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	Category string  `json:"category" db:"category"`
	Cuisine  string  `json:"cuisine" db:"cuisine"`
	Price    float64 `json:"price" db:"price"`
	RestoID  string  `json:"restaurant_id" db:"id"`
}

type MenuItemES struct {
	Name     string  `json:"name" db:"name"`
	Category string  `json:"category" db:"category"`
	Cuisine  string  `json:"cuisine" db:"cuisine"`
	Price    float64 `json:"price" db:"price"`
	RestoID  string  `json:"restaurant_id" db:"id"`
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
