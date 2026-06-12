package handlers

import (
	"encoding/json"
	"net/http"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"

	"github.com/go-chi/chi/v5"
)

// TransporteHandler estructurará las dependencias de la base de datos más adelante
type TransporteHandler struct{}

// NewTransporteHandler crea una nueva instancia del manejador
func NewTransporteHandler() *TransporteHandler {
	return &TransporteHandler{}
}

// Rutas expone todos los endpoints del módulo mediante un Subrouter de Chi
func (h *TransporteHandler) Rutas() http.Handler {
	r := chi.NewRouter()

	// Definición preliminar de los 5 Endpoints CRUD requeridos por la rúbrica
	r.Post("/", h.CrearSolicitud)          // POST /api/v1/transporte
	r.Get("/", h.ObtenerSolicitudes)       // GET /api/v1/transporte
	r.Get("/{id}", h.ObtenerSolicitudPorID) // GET /api/v1/transporte/{id}
	r.Put("/{id}", h.ActualizarSolicitud)   // PUT /api/v1/transporte/{id}
	r.Delete("/{id}", h.EliminarSolicitud) // DELETE /api/v1/transporte/{id}

	return r
}

// --- MOCKS TEMPORALES PARA EVITAR ERRORES DE COMPILACIÓN ---

func (h *TransporteHandler) CrearSolicitud(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mock: Crear Solicitud"))
}

func (h *TransporteHandler) ObtenerSolicitudes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mock: Obtener todas las Solicitudes"))
}

func (h *TransporteHandler) ObtenerSolicitudPorID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mock: Obtener Solicitud por ID"))
}

func (h *TransporteHandler) ActualizarSolicitud(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mock: Actualizar Solicitud"))
}

func (h *TransporteHandler) EliminarSolicitud(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mock: Eliminar Solicitud"))
}