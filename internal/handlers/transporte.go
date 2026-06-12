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

func (h *TransporteHandler) Rutas() http.Handler {
	r := chi.NewRouter()

	// Solicitudes
	r.Route("/solicitudes", func(r chi.Router) {
		r.Get("/", h.ListarSolicitudes)
		r.Get("/{id}", h.ObtenerSolicitud)
		r.Post("/", h.CrearSolicitud)
		r.Put("/{id}", h.ActualizarSolicitud)
		r.Delete("/{id}", h.BorrarSolicitud)
	})

	// Rutas
	r.Route("/rutas", func(r chi.Router) {
		r.Get("/", h.ListarRutas)
		r.Get("/{id}", h.ObtenerRuta)
		r.Post("/", h.CrearRuta)
		r.Put("/{id}", h.ActualizarRuta)
		r.Delete("/{id}", h.BorrarRuta)
	})

	// Paradas
	r.Route("/paradas", func(r chi.Router) {
		r.Get("/", h.ListarParadas)
		r.Post("/", h.CrearParada)
	})

	// Carritos
	r.Route("/carritos", func(r chi.Router) {
		r.Get("/", h.ListarCarritos)
		r.Post("/", h.CrearCarrito)
	})

	// Locaciones
	r.Route("/locaciones", func(r chi.Router) {
		r.Get("/", h.ListarLocaciones)
		r.Post("/", h.CrearLocacion)
	})

	return r
}

// --- Métodos de Rutas ---
func (h *TransporteHandler) CrearRuta(w http.ResponseWriter, r *http.Request) {
	var req modelos.Ruta
	json.NewDecoder(r.Body).Decode(&req)
	ruta := h.store.CrearRuta(req)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ruta)
}

func (h *TransporteHandler) ActualizarRuta(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req modelos.Ruta
	json.NewDecoder(r.Body).Decode(&req)
	ruta, ok := h.store.ActualizarRuta(id, req)
	if !ok { http.Error(w, "no encontrada", 404); return }
	json.NewEncoder(w).Encode(ruta)
}

func (h *TransporteHandler) BorrarRuta(w http.ResponseWriter, r *http.Request) {
	if !h.store.BorrarRuta(chi.URLParam(r, "id")) { http.Error(w, "no encontrada", 404); return }
	w.Write([]byte(`{"mensaje":"eliminado"}`))
}

// --- Métodos de Paradas ---
func (h *TransporteHandler) ListarParadas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.store.ListarParadas())
}
func (h *TransporteHandler) CrearParada(w http.ResponseWriter, r *http.Request) {
	var p modelos.Parada
	json.NewDecoder(r.Body).Decode(&p)
	json.NewEncoder(w).Encode(h.store.CrearParada(p))
}

// --- Métodos de Carritos ---
func (h *TransporteHandler) ListarCarritos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.store.ListarCarritos())
}
func (h *TransporteHandler) CrearCarrito(w http.ResponseWriter, r *http.Request) {
	var c modelos.Carrito
	json.NewDecoder(r.Body).Decode(&c)
	json.NewEncoder(w).Encode(h.store.CrearCarrito(c))
}

// --- Métodos de Locaciones ---
func (h *TransporteHandler) ListarLocaciones(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.store.ListarLocaciones())
}
func (h *TransporteHandler) CrearLocacion(w http.ResponseWriter, r *http.Request) {
	var l modelos.Locacion
	json.NewDecoder(r.Body).Decode(&l)
	json.NewEncoder(w).Encode(h.store.CrearLocacion(l))
}

// ... Mantén los métodos de Solicitudes que ya tenías ...