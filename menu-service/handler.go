package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("menu-service OK"))
}

// RestoHandler handler for Restaurant resource
type RestoHandler struct {
	repo *RestoRepo
}

func NewRestoHandler(repo *RestoRepo) *RestoHandler {
	return &RestoHandler{
		repo: repo,
	}
}

func (h *RestoHandler) NewRestoHandler(w http.ResponseWriter, r *http.Request) {
	resto := &Restaurant{}
	if err := json.NewDecoder(r.Body).Decode(resto); err != nil {
		http.Error(w, fmt.Sprintf("Error: new restaurant json parsing failed, %v", err), http.StatusBadRequest)
		return
	}

	id, err := h.repo.CreateResto(resto)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: new restaurant insert record failed, %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&IDResponse{id})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: restaurant created but restaurant id to json failed, %v", err), http.StatusInternalServerError)
		return
	}
	log.Infof("restaurant created, id=%s", id)
}

func (h *RestoHandler) GetRestoHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := uuid.FromString(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: invalid id '%s'", id), http.StatusBadRequest)
		return
	}

	resto, err := h.repo.GetResto(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: get resto '%s' failed, %v", id, err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(resto)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: restaurant to json failed, %v", err), http.StatusInternalServerError)
		return
	}
}

// MenuItemHandler handler for MenuItem resource
type MenuItemHandler struct {
	repo      *MenuItemRepo
	publisher *Publisher
}

func NewMenuItemHandler(repo *MenuItemRepo, publisher *Publisher) *MenuItemHandler {
	return &MenuItemHandler{
		repo:      repo,
		publisher: publisher,
	}
}

func (h *MenuItemHandler) NewMenuItemHandler(w http.ResponseWriter, r *http.Request) {
	menuItem := &MenuItem{}
	if err := json.NewDecoder(r.Body).Decode(menuItem); err != nil {
		http.Error(w, fmt.Sprintf("Error: new menu item json parsing failed, %v", err), http.StatusBadRequest)
		return
	}

	id, err := h.repo.CreateMenuItem(menuItem)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: new menu item insert record failed, %v", err), http.StatusInternalServerError)
		return
	}
	log.Infof("menuitem created, id=%s", id)
	menuItem.ID = id

	err = h.publisher.PublishMenuItem(menuItem)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: new menu item publish failed, %v", err), http.StatusInternalServerError)
		return
	}
	log.Infof("menuitem published, id=%s", menuItem.ID)

	err = json.NewEncoder(w).Encode(&IDResponse{id})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: menu item created but menu item id to json failed, %v", err), http.StatusInternalServerError)
		return
	}
}
