package handlers

import (
	"encoding/json"
	"net/http"
)

type Server struct {
    Storage storage.Almacen
}

// NewTransporteHandler es el constructor que necesitas en main.go
func NewTransporteHandler(a storage.Almacen) *Server {
    return &Server{Storage: a}
}
// responderJSON escribe una respuesta JSON con el status code dado
func responderJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "error al serializar respuesta", http.StatusInternalServerError)
	}
}

// responderError escribe un mensaje de error en formato JSON
func responderError(w http.ResponseWriter, status int, mensaje string) {
	responderJSON(w, status, map[string]string{"error": mensaje})
}
