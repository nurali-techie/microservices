package main

import (
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// RestoRepo repo for Restaurant entity
type RestoRepo struct {
	db *sql.DB
}

func NewRestoRepo(db *sql.DB) *RestoRepo {
	return &RestoRepo{
		db: db,
	}
}

func (r *RestoRepo) CreateResto(resto *Restaurant) (string, error) {
	restoID := resto.ID
	if restoID == "" {
		restoID = uuid.NewV4().String()
	}

	stmt, err := r.db.Prepare(`INSERT INTO restaurants (id, name, address, city) values ($1, $2, $3, $4) returning id`)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(restoID, resto.Name, resto.Address, resto.City)
	return restoID, err
}

func (r *RestoRepo) GetResto(restoID string) (*Restaurant, error) {
	stmt, err := r.db.Prepare("SELECT id, name, address, city FROM restaurants where id = $1")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(restoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		resto := &Restaurant{}
		err = rows.Scan(&resto.ID, &resto.Name, &resto.Address, &resto.City)
		return resto, err
	}

	return nil, fmt.Errorf("restaurant not found '%s'", restoID)
}

// MenuItemRepo repo for MenuItem entity
type MenuItemRepo struct {
	db *sql.DB
}

func NewMenuItemRepo(db *sql.DB) *MenuItemRepo {
	return &MenuItemRepo{
		db: db,
	}
}

func (r *MenuItemRepo) CreateMenuItem(menuItem *MenuItem) (string, error) {
	miID := menuItem.ID
	if miID == "" {
		miID = uuid.NewV4().String()
	}

	stmt, err := r.db.Prepare(`INSERT INTO menu_items (id, name, category, cuisine, price, restaurant_id) values ($1, $2, $3, $4, $5, $6) returning id`)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(miID, menuItem.Name, menuItem.Category, menuItem.Cuisine, menuItem.Price, menuItem.RestoID)
	return miID, err
}
