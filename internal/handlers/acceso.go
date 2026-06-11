package handlers

import (
	"encoding/json"
	"net/http"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/storage"

	"github.com/go-chi/chi/v5"
)

type AccesoHandler struct {
	Storage *storage.Memoria
}

func NewAccesoHandler(store *storage.Memoria) *AccesoHandler {
	return &AccesoHandler{Storage: store}
}

// Create: POST /api/v1/acceso
func (h *AccesoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var nuevo modelos.Acceso
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Validación básica requerida por la rúbrica
	if nuevo.PlacaVehiculo == "" || nuevo.IDPuntoAcceso <= 0 {
		http.Error(w, "La placa del vehículo y el ID del punto de acceso son obligatorios", http.StatusBadRequest)
		return
	}

	// CORREGIDO: Usa el método CrearAcceso de tu equipo
	resultado := h.Storage.CrearAcceso(nuevo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}

// Read One: GET /api/v1/acceso/{id}
func (h *AccesoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// Llama a BuscarAccesoPorID pasándole el string directamente
	acceso, existe := h.Storage.BuscarAccesoPorID(idStr)
	if !existe {
		http.Error(w, "Registro no encontrado", http.StatusNotFound) // 404 Not Found [cite: 38, 44]
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acceso)
}

// Read All: GET /api/v1/acceso
func (h *AccesoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// Usa el método ListarAccesos de tu equipo
	accesos := h.Storage.ListarAccesos()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accesos)
}

// Update: PUT /api/v1/acceso/{id}
func (h *AccesoHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	var datosActualizados modelos.Acceso
	if err := json.NewDecoder(r.Body).Decode(&datosActualizados); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Usa el método ActualizarAcceso de tu equipo y verifica si existe
	accesoModificado, existe := h.Storage.ActualizarAcceso(idStr, datosActualizados)
	if !existe {
		http.Error(w, "Registro no encontrado para actualizar", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accesoModificado)
}

// Delete: DELETE /api/v1/acceso/{id}
func (h *AccesoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// Usa el método EliminarAcceso de tu equipo y verifica el booleano
	eliminado := h.Storage.EliminarAcceso(idStr)
	if !eliminado {
		http.Error(w, "El registro no existe", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"mensaje": "Registro eliminado exitosamente"}`))
}
