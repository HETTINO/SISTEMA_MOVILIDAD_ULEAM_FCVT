package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/storage"

	"github.com/go-chi/chi/v5"
)



type TransporteHandler struct {
	Storage *storage.Memoria
}

func NewTransporteHandler(store *storage.Memoria) *TransporteHandler {
	return &TransporteHandler{Storage: store}
}

// ==========================================
// CONTROLADORES: SOLICITUDES
// ==========================================
func (h *TransporteHandler) ListarSolicitudes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Storage.ListarSolicitudes())
}

func (h *TransporteHandler) ObtenerSolicitud(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido numérico requerido", http.StatusBadRequest)
		return
	}

	res, existe := h.Storage.BuscarSolicitudPorID(id)
	if !existe {
		http.Error(w, "Solicitud no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *TransporteHandler) CrearSolicitud(w http.ResponseWriter, r *http.Request) {
	var s modelos.Solicitud
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Validación básica requerida por la rúbrica
	if s.CedulaUsuario == "" || s.PuntoDestino == "" || s.CantPersonas <= 0 {
		http.Error(w, "Cédula, destino y cantidad de personas válidas son requeridos", http.StatusBadRequest)
		return
	}

	s.Estado = "Pendiente"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(h.Storage.CrearSolicitud(s))
}

func (h *TransporteHandler) ActualizarSolicitud(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var datos modelos.Solicitud
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	modificado, existe := h.Storage.ActualizarSolicitud(id, datos)
	if !existe {
		http.Error(w, "Solicitud no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modificado)
}

func (h *TransporteHandler) EliminarSolicitud(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !h.Storage.EliminarSolicitud(id) {
		http.Error(w, "La solicitud no existe", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"mensaje": "Solicitud eliminada exitosamente"}`))
}

// ==========================================
// CONTROLADORES: CARRITOS
// ==========================================
func (h *TransporteHandler) ListarCarritos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Storage.ListarCarritos())
}

func (h *TransporteHandler) ObtenerCarrito(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	res, existe := h.Storage.BuscarCarritoPorID(id)
	if !existe {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *TransporteHandler) CrearCarrito(w http.ResponseWriter, r *http.Request) {
	var c modelos.Carrito
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if c.NombreCarrito == "" || c.Capacidad <= 0 {
		http.Error(w, "Nombre del carrito y capacidad válida son requeridos", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(h.Storage.CrearCarrito(c))
}

func (h *TransporteHandler) ActualizarCarrito(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var datos modelos.Carrito
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	modificado, existe := h.Storage.ActualizarCarrito(id, datos)
	if !existe {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modificado)
}

func (h *TransporteHandler) EliminarCarrito(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !h.Storage.EliminarCarrito(id) {
		http.Error(w, "El carrito no existe", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"mensaje": "Carrito eliminado exitosamente"}`))
}

// ==========================================
// CONTROLADORES: LOCACIONES
// ==========================================
func (h *TransporteHandler) ListarLocaciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Storage.ListarLocaciones())
}

func (h *TransporteHandler) ObtenerLocacion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	res, existe := h.Storage.BuscarLocacionPorID(id)
	if !existe {
		http.Error(w, "Locación no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *TransporteHandler) CrearLocacion(w http.ResponseWriter, r *http.Request) {
	var l modelos.Locacion
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if l.IDCarrito <= 0 || l.Latitud == 0 || l.Longitud == 0 {
		http.Error(w, "IDCarrito, Latitud y Longitud son requeridos", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(h.Storage.CrearLocacion(l))
}

func (h *TransporteHandler) ActualizarLocacion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var datos modelos.Locacion
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	modificado, existe := h.Storage.ActualizarLocacion(id, datos)
	if !existe {
		http.Error(w, "Locación no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modificado)
}

func (h *TransporteHandler) EliminarLocacion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !h.Storage.EliminarLocacion(id) {
		http.Error(w, "La locación no existe", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"mensaje": "Locación eliminada exitosamente"}`))
}

// ==========================================
// CONTROLADORES: RUTAS
// ==========================================
func (h *TransporteHandler) ListarRutas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Storage.ListarRutas())
}

func (h *TransporteHandler) ObtenerRuta(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	res, existe := h.Storage.BuscarRutaPorID(id)
	if !existe {
		http.Error(w, "Ruta no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *TransporteHandler) CrearRuta(w http.ResponseWriter, r *http.Request) {
	var rt modelos.Ruta
	if err := json.NewDecoder(r.Body).Decode(&rt); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if rt.NombreRuta == "" {
		http.Error(w, "El nombre de la ruta es requerido", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(h.Storage.CrearRuta(rt))
}

func (h *TransporteHandler) ActualizarRuta(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var datos modelos.Ruta
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	modificado, existe := h.Storage.ActualizarRuta(id, datos)
	if !existe {
		http.Error(w, "Ruta no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modificado)
}

func (h *TransporteHandler) EliminarRuta(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !h.Storage.EliminarRuta(id) {
		http.Error(w, "La ruta no existe", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"mensaje": "Ruta eliminada exitosamente"}`))
}

// ==========================================
// CONTROLADORES: PARADAS
// ==========================================
func (h *TransporteHandler) ListarParadas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Storage.ListarParadas())
}

func (h *TransporteHandler) ObtenerParada(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	res, existe := h.Storage.BuscarParadaPorID(id)
	if !existe {
		http.Error(w, "Parada no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *TransporteHandler) CrearParada(w http.ResponseWriter, r *http.Request) {
	var p modelos.Paradas
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if p.Nombre == "" || p.IDRuta <= 0 {
		http.Error(w, "El nombre de la parada y un IDRuta válido son requeridos", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(h.Storage.CrearParada(p))
}

func (h *TransporteHandler) ActualizarParada(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var datos modelos.Paradas
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	modificado, existe := h.Storage.ActualizarParada(id, datos)
	if !existe {
		http.Error(w, "Parada no encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modificado)
}

func (h *TransporteHandler) EliminarParada(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !h.Storage.EliminarParada(id) {
		http.Error(w, "La parada no existe", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"mensaje": "Parada eliminada exitosamente"}`))
}