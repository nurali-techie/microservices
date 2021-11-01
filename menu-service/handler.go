package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("menu-service OK"))
}

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
