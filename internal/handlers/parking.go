package handlers

import (
	"encoding/json"
	"net/http"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/storage"

	"github.com/go-chi/chi/v5"
)

type ParkingHandler struct {
	store storage.Almacen
}

func NewParkingHandler(store storage.Almacen) *ParkingHandler {
	return &ParkingHandler{store: store}
}

// GET /api/v1/parking
func (h *ParkingHandler) ListarParqueaderos(w http.ResponseWriter, r *http.Request) {
	parqueaderos := h.store.ListarParqueaderos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parqueaderos)
}

// GET /api/v1/parking/{id}
func (h *ParkingHandler) ObtenerParqueadero(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parqueadero, ok := h.store.BuscarParqueaderoporID(id)
	if !ok {
		http.Error(w, "parqueadero no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parqueadero)
}

// POST /api/v1/parking
func (h *ParkingHandler) CrearParqueadero(w http.ResponseWriter, r *http.Request) {
	var req modelos.CrearParqueaderoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	// Validación básica
	if req.Nombre == "" || req.Ubicacion == "" || req.Capacidad <= 0 {
		http.Error(w, "nombre, ubicacion y capacidad son requeridos", http.StatusBadRequest)
		return
	}
	parqueadero := h.store.CrearParqueadero(req)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(parqueadero)
}

// PUT /api/v1/parking/{id}
func (h *ParkingHandler) ActualizarParqueadero(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req modelos.ActualizarParqueaderoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	parqueadero, ok := h.store.ActualizarParqueadero(id, req)
	if !ok {
		http.Error(w, "parqueadero no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parqueadero)
}

// DELETE /api/v1/parking/{id}
func (h *ParkingHandler) EliminarParqueadero(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if !h.store.EliminarParqueadero(id) {
		http.Error(w, "parqueadero no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "parqueadero eliminado correctamente"})
}
