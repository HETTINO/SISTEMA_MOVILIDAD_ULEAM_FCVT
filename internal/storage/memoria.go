package storage

import (
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"
	"strconv"
	"sync"
	"time"
)

type Memoria struct {
	mu sync.RWMutex

	parking     []modelos.Parqueadero
	espacios    []modelos.Espacio
	ocupaciones []modelos.Ocupacion
	rutas       []modelos.Ruta
	paradas     []modelos.Parada
	solicitudes []modelos.Solicitud
	carritos    []modelos.Carrito
	locaciones  []modelos.Locacion

	nextparkingID     int
	nextEspacioID     int
	nextOcupacionID   int
	nextRutaID        int
	nextSolicitudID   int
	nextParadaID      int
	nextCarritoID     int
	nextLocacionID    int
}

func NewMemoria() *Memoria {
	return &Memoria{
		parking: []modelos.Parqueadero{}, espacios: []modelos.Espacio{}, ocupaciones: []modelos.Ocupacion{},
		rutas: []modelos.Ruta{}, paradas: []modelos.Parada{}, solicitudes: []modelos.Solicitud{},
		carritos: []modelos.Carrito{}, locaciones: []modelos.Locacion{},
		nextparkingID: 1, nextEspacioID: 1, nextOcupacionID: 1, nextRutaID: 1,
		nextSolicitudID: 1, nextParadaID: 1, nextCarritoID: 1, nextLocacionID: 1,
	}
}

// --- MÉTODOS SOLICITUDES ---
func (m *Memoria) ListarSolicitudes() []modelos.Solicitud {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c := make([]modelos.Solicitud, len(m.solicitudes))
	copy(c, m.solicitudes)
	return c
}

func (m *Memoria) BuscarSolicitudPorID(id string) (modelos.Solicitud, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, s := range m.solicitudes {
		if strconv.Itoa(s.ID) == id { return s, true }
	}
	return modelos.Solicitud{}, false
}

func (m *Memoria) CrearSolicitud(req modelos.Solicitud) modelos.Solicitud {
	m.mu.Lock()
	defer m.mu.Unlock()
	req.ID = m.nextSolicitudID
	m.nextSolicitudID++
	m.solicitudes = append(m.solicitudes, req)
	return req
}

func (m *Memoria) ActualizarSolicitud(id string, req modelos.Solicitud) (modelos.Solicitud, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.solicitudes {
		if strconv.Itoa(s.ID) == id {
			m.solicitudes[i].CedulaUsuario = req.CedulaUsuario
			m.solicitudes[i].CantPersonas = req.CantPersonas
			m.solicitudes[i].PuntoDestino = req.PuntoDestino
			m.solicitudes[i].Estado = req.Estado
			m.solicitudes[i].IDCarrito = req.IDCarrito
			return m.solicitudes[i], true
		}
	}
	return modelos.Solicitud{}, false
}

func (m *Memoria) BorrarSolicitud(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.solicitudes {
		if strconv.Itoa(s.ID) == id {
			m.solicitudes = append(m.solicitudes[:i], m.solicitudes[i+1:]...)
			return true
		}
	}
	return false
}

// --- MÉTODOS RUTAS ---
func (m *Memoria) ListarRutas() []modelos.Ruta {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c := make([]modelos.Ruta, len(m.rutas))
	copy(c, m.rutas)
	return c
}

func (m *Memoria) BuscarRutaPorID(id string) (modelos.Ruta, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, r := range m.rutas {
		if strconv.Itoa(r.ID) == id { return r, true }
	}
	return modelos.Ruta{}, false
}

func (m *Memoria) CrearRuta(r modelos.Ruta) modelos.Ruta {
	m.mu.Lock()
	defer m.mu.Unlock()
	r.ID = m.nextRutaID
	m.nextRutaID++
	m.rutas = append(m.rutas, r)
	return r
}

func (m *Memoria) ActualizarRuta(id string, req modelos.Ruta) (modelos.Ruta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, r := range m.rutas {
		if strconv.Itoa(r.ID) == id {
			m.rutas[i].Nombre = req.Nombre
			m.rutas[i].Descripcion = req.Descripcion
			return m.rutas[i], true
		}
	}
	return modelos.Ruta{}, false
}

func (m *Memoria) BorrarRuta(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, r := range m.rutas {
		if strconv.Itoa(r.ID) == id {
			m.rutas = append(m.rutas[:i], m.rutas[i+1:]...)
			return true
		}
	}
	return false
}

// --- MÉTODOS PARADAS ---
func (m *Memoria) ListarParadas() []modelos.Parada {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c := make([]modelos.Parada, len(m.paradas))
	copy(c, m.paradas)
	return c
}

func (m *Memoria) BuscarParadaPorID(id string) (modelos.Parada, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, p := range m.paradas {
		if strconv.Itoa(p.IDParada) == id { return p, true }
	}
	return modelos.Parada{}, false
}

func (m *Memoria) CrearParada(p modelos.Parada) modelos.Parada {
	m.mu.Lock()
	defer m.mu.Unlock()
	p.IDParada = m.nextParadaID
	m.nextParadaID++
	m.paradas = append(m.paradas, p)
	return p
}

func (m *Memoria) ActualizarParada(id string, datos modelos.Parada) (modelos.Parada, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.paradas {
		if strconv.Itoa(p.IDParada) == id {
			m.paradas[i].Nombre = datos.Nombre
			return m.paradas[i], true
		}
	}
	return modelos.Parada{}, false
}

func (m *Memoria) BorrarParada(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.paradas {
		if strconv.Itoa(p.IDParada) == id {
			m.paradas = append(m.paradas[:i], m.paradas[i+1:]...)
			return true
		}
	}
	return false
}

// --- MÉTODOS CARRITOS ---
func (m *Memoria) ListarCarritos() []modelos.Carrito {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c := make([]modelos.Carrito, len(m.carritos))
	copy(c, m.carritos)
	return c
}

func (m *Memoria) BuscarCarritoPorID(id string) (modelos.Carrito, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, c := range m.carritos {
		if strconv.Itoa(c.ID) == id { return c, true }
	}
	return modelos.Carrito{}, false
}

func (m *Memoria) CrearCarrito(c modelos.Carrito) modelos.Carrito {
	m.mu.Lock()
	defer m.mu.Unlock()
	c.ID = m.nextCarritoID
	m.nextCarritoID++
	m.carritos = append(m.carritos, c)
	return c
}

func (m *Memoria) ActualizarCarrito(id string, datos modelos.Carrito) (modelos.Carrito, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.carritos {
		if strconv.Itoa(c.ID) == id {
			m.carritos[i].NombreCarrito = datos.NombreCarrito
			m.carritos[i].Capacidad = datos.Capacidad
			m.carritos[i].Estado = datos.Estado
			m.carritos[i].RutaID = datos.RutaID
			return m.carritos[i], true
		}
	}
	return modelos.Carrito{}, false
}

func (m *Memoria) BorrarCarrito(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.carritos {
		if strconv.Itoa(c.ID) == id {
			m.carritos = append(m.carritos[:i], m.carritos[i+1:]...)
			return true
		}
	}
	return false
}

// --- MÉTODOS LOCACIONES ---
func (m *Memoria) ListarLocaciones() []modelos.Locacion {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c := make([]modelos.Locacion, len(m.locaciones))
	copy(c, m.locaciones)