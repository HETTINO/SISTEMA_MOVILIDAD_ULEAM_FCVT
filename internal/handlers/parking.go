package handlers

import (
	"encoding/json"
	"net/http"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/storage"

	"github.com/go-chi/chi/v5"
)

// ParkingHandler agrupa todos los handlers del módulo de parqueo
type ParkingHandler struct {
	store *storage.ParkingStore
}

// NewParkingHandler crea un handler con su store
func NewParkingHandler(store *storage.ParkingStore) *ParkingHandler {
	return &ParkingHandler{store: store}
}

// =============================================
// PARQUEADEROS CRUD
// =============================================

// ListarParqueaderos GET /api/v1/parking
// Devuelve todos los parqueaderos registrados
func (h *ParkingHandler) ListarParqueaderos(w http.ResponseWriter, r *http.Request) {
	lista := h.store.ListarParqueaderos()
	responderJSON(w, http.StatusOK, lista)
}

// ObtenerParqueadero GET /api/v1/parking/{id}
// Devuelve un parqueadero por su ID
func (h *ParkingHandler) ObtenerParqueadero(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	p, err := h.store.ObtenerParqueadero(id)
	if err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}
	responderJSON(w, http.StatusOK, p)
}

// CrearParqueadero POST /api/v1/parking
// Crea un nuevo parqueadero
func (h *ParkingHandler) CrearParqueadero(w http.ResponseWriter, r *http.Request) {
	var req modelos.CrearParqueaderoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responderError(w, http.StatusBadRequest, "cuerpo de la solicitud inválido")
		return
	}
	p, err := h.store.CrearParqueadero(req)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}
	responderJSON(w, http.StatusCreated, p)
}

// ActualizarParqueadero PUT /api/v1/parking/{id}
// Actualiza los datos de un parqueadero existente
func (h *ParkingHandler) ActualizarParqueadero(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req modelos.CrearParqueaderoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responderError(w, http.StatusBadRequest, "cuerpo de la solicitud inválido")
		return
	}
	p, err := h.store.ActualizarParqueadero(id, req)
	if err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}
	responderJSON(w, http.StatusOK, p)
}

// EliminarParqueadero DELETE /api/v1/parking/{id}
// Elimina un parqueadero por su ID
func (h *ParkingHandler) EliminarParqueadero(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.store.EliminarParqueadero(id); err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}
	responderJSON(w, http.StatusOK, map[string]string{"mensaje": "parqueadero eliminado"})
}

// =============================================
// ESPACIOS CRUD
// =============================================

// ListarEspacios GET /api/v1/parking/{id}/espacios
// Devuelve los espacios de un parqueadero
func (h *ParkingHandler) ListarEspacios(w http.ResponseWriter, r *http.Request) {
	parqueaderoID := chi.URLParam(r, "id")
	espacios := h.store.ListarEspacios(parqueaderoID)
	responderJSON(w, http.StatusOK, espacios)
}

// ObtenerEspacio GET /api/v1/parking/espacios/{espacioID}
// Devuelve un espacio por su ID
func (h *ParkingHandler) ObtenerEspacio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "espacioID")
	e, err := h.store.ObtenerEspacio(id)
	if err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}
	responderJSON(w, http.StatusOK, e)
}

// CrearEspacio POST /api/v1/parking/espacios
// Crea un nuevo espacio dentro de un parqueadero
func (h *ParkingHandler) CrearEspacio(w http.ResponseWriter, r *http.Request) {
	var req modelos.CrearEspacioRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responderError(w, http.StatusBadRequest, "cuerpo de la solicitud inválido")
		return
	}
	e, err := h.store.CrearEspacio(req)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}
	responderJSON(w, http.StatusCreated, e)
}

// ReservarEspacio PATCH /api/v1/parking/espacios/{espacioID}/reserve
// Marca un espacio como ocupado (reserva por 5 minutos)
func (h *ParkingHandler) ReservarEspacio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "espacioID")
	e, err := h.store.ReservarEspacio(id)
	if err != nil {
		// Si no se encuentra → 404, si ya está ocupado → 400
		status := http.StatusBadRequest
		if err.Error() == "espacio no encontrado" {
			status = http.StatusNotFound
		}
		responderError(w, status, err.Error())
		return
	}
	responderJSON(w, http.StatusOK, e)
}

// LiberarEspacio PATCH /api/v1/parking/espacios/{espacioID}/release
// Marca un espacio como disponible
func (h *ParkingHandler) LiberarEspacio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "espacioID")
	e, err := h.store.LiberarEspacio(id)
	if err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}
	responderJSON(w, http.StatusOK, e)
}

// EliminarEspacio DELETE /api/v1/parking/espacios/{espacioID}
// Elimina un espacio del sistema
func (h *ParkingHandler) EliminarEspacio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "espacioID")
	if err := h.store.EliminarEspacio(id); err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}
	responderJSON(w, http.StatusOK, map[string]string{"mensaje": "espacio eliminado"})
}

// =============================================
// DISPONIBILIDAD Y OCUPACIÓN
// =============================================

// ObtenerDisponibilidad GET /api/v1/parking/{id}/disponibilidad
// Devuelve cuántos espacios libres y ocupados hay en un parqueadero
func (h *ParkingHandler) ObtenerDisponibilidad(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	disp, err := h.store.ObtenerDisponibilidad(id)
	if err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}
	responderJSON(w, http.StatusOK, disp)
}

// ObtenerEstado GET /api/v1/parking/{id}/status
// Devuelve el estado general + lista completa de espacios
func (h *ParkingHandler) ObtenerEstado(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	disp, err := h.store.ObtenerDisponibilidad(id)
	if err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}
	espacios := h.store.ListarEspacios(id)
	respuesta := map[string]any{
		"resumen":  disp,
		"espacios": espacios,
	}
	responderJSON(w, http.StatusOK, respuesta)
}

// ListarOcupaciones GET /api/v1/parking/ocupacion
// Devuelve todas las ocupaciones activas
func (h *ParkingHandler) ListarOcupaciones(w http.ResponseWriter, r *http.Request) {
	ocupaciones := h.store.ListarOcupaciones()
	responderJSON(w, http.StatusOK, ocupaciones)
}

// RegistrarOcupacion POST /api/v1/parking/espacios/{espacioID}/ocupar
// Registra la entrada de un vehículo a un espacio
func (h *ParkingHandler) RegistrarOcupacion(w http.ResponseWriter, r *http.Request) {
	espacioID := chi.URLParam(r, "espacioID")
	var req modelos.OcuparEspacioRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responderError(w, http.StatusBadRequest, "cuerpo de la solicitud inválido")
		return
	}
	o, err := h.store.RegistrarOcupacion(req, espacioID)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "espacio no encontrado" {
			status = http.StatusNotFound
		}
		responderError(w, status, err.Error())
		return
	}
	responderJSON(w, http.StatusCreated, o)
}
