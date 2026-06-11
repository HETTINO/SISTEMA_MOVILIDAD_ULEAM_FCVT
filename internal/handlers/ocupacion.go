package handlers

import (
	"encoding/json"
	"net/http"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"

	"github.com/go-chi/chi/v5"
)

// GET /api/v1/ocupaciones
func (h *ParkingHandler) ListarOcupaciones(w http.ResponseWriter, r *http.Request) {

	ocupaciones := h.store.ListarOcupaciones()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ocupaciones)
}

// GET /api/v1/ocupaciones/{id}
func (h *ParkingHandler) ObtenerOcupacion(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	ocupacion, ok := h.store.BuscarOcupacionPorID(id)

	if !ok {
		http.Error(w, "Ocupación no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ocupacion)
}

// POST /api/v1/ocupaciones
func (h *ParkingHandler) CrearOcupacion(w http.ResponseWriter, r *http.Request) {

	var req modelos.OcuparEspacioRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	ocupacion, ok := h.store.CrearOcupacion(req)

	if !ok {
		http.Error(w, "No se pudo registrar la ocupación", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(ocupacion)
}

// PUT /api/v1/ocupaciones/{id}/liberar
func (h *ParkingHandler) LiberarOcupacion(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	ocupacion, ok := h.store.LiberarOcupacion(id)

	if !ok {
		http.Error(w, "Ocupación no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ocupacion)
}
