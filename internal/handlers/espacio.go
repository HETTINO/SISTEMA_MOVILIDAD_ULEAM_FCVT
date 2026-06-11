package handlers

import (
	"encoding/json"
	"net/http"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"

	"github.com/go-chi/chi/v5"
)

// GET /api/v1/espacios
func (h *ParkingHandler) ListarEspacios(w http.ResponseWriter, r *http.Request) {
	espacios := h.store.ListarEspacios()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(espacios)
}

// GET /api/v1/espacios/{id}
func (h *ParkingHandler) ObtenerEspacio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	espacio, ok := h.store.BuscarEspacioPorID(id)
	if !ok {
		http.Error(w, "Espacio no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(espacio)
}

// POST /api/v1/espacios
func (h *ParkingHandler) CrearEspacio(w http.ResponseWriter, r *http.Request) {
	var req modelos.CrearEspacioRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	espacio, ok := h.store.CrearEspacio(req)
	if !ok {
		http.Error(w, "No se pudo crear el espacio", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(espacio)
}

// PUT /api/v1/espacios/{id}
func (h *ParkingHandler) ActualizarEspacio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req modelos.ActualizarEspacioRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	espacio, ok := h.store.ActualizarEspacio(id, req)
	if !ok {
		http.Error(w, "Espacio no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(espacio)
}

// DELETE /api/v1/espacios/{id}
func (h *ParkingHandler) EliminarEspacio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if !h.store.EliminarEspacio(id) {
		http.Error(w, "Espacio no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Espacio eliminado correctamente",
		"id":      id,
	})
}
