package main

import (
	"encoding/json"
	"net/http"
)

//Usersdb usuarios de base de datos
func Usersdb(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	a, _ := obtenerResultadosdb()
	json.NewEncoder(w).Encode(a)
}
