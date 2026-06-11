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

// ==========================================
// CONTROLADORES: USUARIOS
// ==========================================
func (h *AccesoHandler) ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Storage.ListarUsuarios())
}

func (h *AccesoHandler) ObtenerUsuario(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, existe := h.Storage.BuscarUsuarioPorID(id)
	if !existe {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *AccesoHandler) CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var u modelos.Usuario
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(h.Storage.CrearUsuario(u))
}

func (h *AccesoHandler) ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (h *AccesoHandler) EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// ==========================================
// CONTROLADORES: VEHÍCULOS
// ==========================================
func (h *AccesoHandler) ListarVehiculos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Storage.ListarVehiculos())
}

func (h *AccesoHandler) ObtenerVehiculo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, existe := h.Storage.BuscarVehiculoPorID(id)
	if !existe {
		http.Error(w, "Vehículo no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *AccesoHandler) CrearVehiculo(w http.ResponseWriter, r *http.Request) {
	var v modelos.Vehiculo
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(h.Storage.CrearVehiculo(v))
}

func (h *AccesoHandler) ActualizarVehiculo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (h *AccesoHandler) EliminarVehiculo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// ==========================================
// CONTROLADORES: PUNTOS DE ACCESO
// ==========================================
func (h *AccesoHandler) ListarPuntosAcceso(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Storage.ListarPuntosAcceso())
}

func (h *AccesoHandler) ObtenerPuntoAcceso(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, existe := h.Storage.BuscarPuntoAccesoPorID(id)
	if !existe {
		http.Error(w, "Punto de acceso no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *AccesoHandler) CrearPuntoAcceso(w http.ResponseWriter, r *http.Request) {
	var p modelos.PuntoDeAcceso
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(h.Storage.CrearPuntoAcceso(p))
}

func (h *AccesoHandler) ActualizarPuntoAcceso(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (h *AccesoHandler) EliminarPuntoAcceso(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// ==========================================
// CONTROLADORES: ACCESOS (TRANSACCIONAL)
// ==========================================
func (h *AccesoHandler) CrearAcceso(w http.ResponseWriter, r *http.Request) {
	var nuevo modelos.Acceso
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	if nuevo.PlacaVehiculo == "" || nuevo.IDPuntoAcceso <= 0 {
		http.Error(w, "La placa y el ID del punto de acceso son requeridos", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(h.Storage.CrearAcceso(nuevo))
}

func (h *AccesoHandler) ObtenerAcceso(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	acceso, existe := h.Storage.BuscarAccesoPorID(idStr)
	if !existe {
		http.Error(w, "Registro no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acceso)
}

func (h *AccesoHandler) ListarAccesos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Storage.ListarAccesos())
}

func (h *AccesoHandler) ActualizarAcceso(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	var datos modelos.Acceso
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	modificado, existe := h.Storage.ActualizarAcceso(idStr, datos)
	if !existe {
		http.Error(w, "Registro no encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modificado)
}

func (h *AccesoHandler) EliminarAcceso(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if !h.Storage.EliminarAcceso(idStr) {
		http.Error(w, "El registro no existe", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"mensaje": "Registro eliminado exitosamente"}`))
}
