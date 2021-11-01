package main

import "net/http"

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("menu-service OK"))
}
