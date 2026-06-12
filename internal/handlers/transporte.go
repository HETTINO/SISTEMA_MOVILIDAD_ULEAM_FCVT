package handlers

import (
	"encoding/json"
	"net/http"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/storage"

	"github.com/go-chi/chi/v5"
)

type TransporteHandler struct {
	store storage.Almacen
}

func NewTransporteHandler(store storage.Almacen) *TransporteHandler {
	return &TransporteHandler{store: store}
}

// Rutas expone todos los endpoints del módulo mediante un Subrouter de Chi
func (h *TransporteHandler) Rutas() http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.ListarSolicitudes)
	r.Get("/{id}", h.ObtenerSolicitud)
	r.Post("/", h.CrearSolicitud)
	r.Put("/{id}", h.ActualizarSolicitud)
	r.Delete("/{id}", h.EliminarSolicitud)

	return r
}

// GET /api/v1/transporte
func (h *TransporteHandler) ListarSolicitudes(w http.ResponseWriter, r *http.Request) {
	solicitudes := h.store.ListarSolicitudes()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solicitudes)
}

// GET /api/v1/transporte/{id}
func (h *TransporteHandler) ObtenerSolicitud(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	solicitud, ok := h.store.BuscarSolicitudPorID(id)
	if !ok {
		http.Error(w, "solicitud no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solicitud)
}

// POST /api/v1/transporte
func (h *TransporteHandler) CrearSolicitud(w http.ResponseWriter, r *http.Request) {
	var req modelos.Solicitud
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Para CrearSolicitud (Línea 62):
	if req.CedulaUsuario == "" || req.CantPersonas <= 0 || req.PuntoDestino == "" {
    http.Error(w, "cedula, cantidad de personas y punto de destino son requeridos", http.StatusBadRequest)
    return
	}

	
	solicitud := h.store.CrearSolicitud(req)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(solicitud)
}

// PUT /api/v1/transporte/{id}
func (h *TransporteHandler) ActualizarSolicitud(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req modelos.Solicitud
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Para ActualizarSolicitud (Línea 82):
	if req.CedulaUsuario == "" || req.CantPersonas <= 0 || req.PuntoDestino == "" {
    http.Error(w, "datos de actualizacion invalidos", http.StatusBadRequest)
    return
}

	solicitud, ok := h.store.ActualizarSolicitud(id, req)
	if !ok {
		http.Error(w, "solicitud no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solicitud)
}

// DELETE /api/v1/transporte/{id}
func (h *TransporteHandler) EliminarSolicitud(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	// Corregido: Usamos BorrarSolicitud para acoplarnos exactamente a la interfaz
	if !h.store.BorrarSolicitud(id) {
		http.Error(w, "solicitud no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "solicitud eliminada correctamente"})
}