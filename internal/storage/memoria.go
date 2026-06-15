package storage

import (
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"
	"strconv"
	"sync"
	"time"
)

type Memoria struct {
	parking     []modelos.Parqueadero
	espacios    []modelos.Espacio
	ocupaciones []modelos.Ocupacion

	nextparkingID   int
	nextEspacioID   int
	nextOcupacionID int

	// === CAMPOS DEL MÓDULO: TRANSPORTE INTERNO ===
	solicitudes []modelos.Solicitud
	carritos    []modelos.Carrito
	locaciones  []modelos.Locacion
	rutas       []modelos.Ruta
	paradas     []modelos.Paradas

	nextSolicitudID int
	nextCarritoID   int
	nextLocacionID  int
	nextRutaID      int
	nextParadaID    int
	// =============================================

	mu sync.RWMutex
}

func NewMemoria() *Memoria {
	return &Memoria{
		parking:     []modelos.Parqueadero{},
		espacios:    []modelos.Espacio{},
		ocupaciones: []modelos.Ocupacion{},

		nextEspacioID:   1,
		nextparkingID:   1,
		nextOcupacionID: 1,

		// === INICIALIZACIÓN DEL MÓDULO TRANSPORTE INTERNO ===
		solicitudes: []modelos.Solicitud{},
		carritos:    []modelos.Carrito{},
		locaciones:  []modelos.Locacion{},
		rutas:       []modelos.Ruta{},
		paradas:     []modelos.Paradas{},

		nextSolicitudID: 1,
		nextCarritoID:   1,
		nextLocacionID:  1,
		nextRutaID:      1,
		nextParadaID:    1,
		// ====================================================
	}
}

func (m *Memoria) SeedParqueaderos() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.parking = []modelos.Parqueadero{
		{ID: "1", Nombre: "Parqueadero FCVT", Ubicacion: "Facultad de Ciencias de la Vida y Tecnologías", Capacidad: 20, Activo: true},
		{ID: "2", Nombre: "Parqueadero Central", Ubicacion: "Centro de la ciudad", Capacidad: 50, Activo: true},
		{ID: "3", Nombre: "Parqueadero Norte", Ubicacion: "Zona norte de la ciudad", Capacidad: 30, Activo: true},
		{ID: "4", Nombre: "Parqueadero Sur", Ubicacion: "Zona sur de la ciudad", Capacidad: 25, Activo: true},
		{ID: "5", Nombre: "Parqueadero Este", Ubicacion: "Zona este de la ciudad", Capacidad: 15, Activo: true},
	}
	m.nextparkingID = 6
}

func (m *Memoria) SeedEspacios() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.espacios = []modelos.Espacio{
		{ID: "1", ParqueaderoID: "1", Numero: "A1", Disponible: true},
		{ID: "2", ParqueaderoID: "1", Numero: "A2", Disponible: true},
		{ID: "3", ParqueaderoID: "1", Numero: "A3", Disponible: true},
		{ID: "4", ParqueaderoID: "2", Numero: "B1", Disponible: true},
		{ID: "5", ParqueaderoID: "2", Numero: "B2", Disponible: true},
	}
	m.nextEspacioID = 6
}

func (m *Memoria) SeedOcupaciones() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ocupaciones = []modelos.Ocupacion{
		{ID: 1, EspacioID: "1", Placa: "abc1234", Entrada: time.Now().Add(-2 * time.Hour), Activa: true},
		{ID: 2, EspacioID: "2", Placa: "xyz5678", Entrada: time.Now().Add(-1 * time.Hour), Activa: true},
		{ID: 3, EspacioID: "3", Placa: "del9012", Entrada: time.Now().Add(-30 * time.Minute), Activa: true},
	}
	m.nextOcupacionID = 4
}

func (m *Memoria) SeedTransporteInterno() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.rutas = []modelos.Ruta{
		{ID: 1, NombreRuta: "Circuito Tecnológico", Descripcion: "Pasa por las facultades de Ingeniería y FCVT"},
		{ID: 2, NombreRuta: "Circuito Central", Descripcion: "Recorrido por administración y canchas centrales"},
	}
	m.nextRutaID = 3

	m.paradas = []modelos.Paradas{
		{ID: 1, IDRuta: 1, Nombre: "Parada FCVT", Latitud: -0.9512, Longitud: -80.7412},
		{ID: 2, IDRuta: 1, Nombre: "Parada Ingeniería", Latitud: -0.9525, Longitud: -80.7435},
	}
	m.nextParadaID = 3

	m.carritos = []modelos.Carrito{
		{ID: 1, IDRuta: 1, NombreCarrito: "Eco-Buga 01", Capacidad: 8, Estado: "Disponible"},
		{ID: 2, IDRuta: 2, NombreCarrito: "Eco-Buga 02", Capacidad: 12, Estado: "En Viaje"},
	}
	m.nextCarritoID = 3

	m.solicitudes = []modelos.Solicitud{
		{ID: 1, CedulaUsuario: "1312345678", IDCarrito: 1, CantPersonas: 3, PuntoDestino: "Parada Ingeniería", Estado: "Pendiente"},
	}
	m.nextSolicitudID = 2

	m.locaciones = []modelos.Locacion{
		{ID: 1, IDCarrito: 1, Latitud: -0.9515, Longitud: -80.7418, TimeStamp: time.Now()},
	}
	m.nextLocacionID = 2
}

// -----Parqueaderos-----
func (m *Memoria) ListarParqueaderos() []modelos.Parqueadero {
	m.mu.Lock()
	defer m.mu.Unlock()
	copias := make([]modelos.Parqueadero, len(m.parking))
	copy(copias, m.parking)
	return copias
}

func (m *Memoria) BuscarParqueaderoporID(id string) (modelos.Parqueadero, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, p := range m.parking {
		if p.ID == id {
			return p, true
		}
	}
	return modelos.Parqueadero{}, false
}

func (m *Memoria) CrearParqueadero(req modelos.CrearParqueaderoRequest) modelos.Parqueadero {
	m.mu.Lock()
	defer m.mu.Unlock()

	p := modelos.Parqueadero{
		ID:        strconv.Itoa(m.nextparkingID),
		Nombre:    req.Nombre,
		Ubicacion: req.Ubicacion,
		Capacidad: req.Capacidad,
		Activo:    true,
	}

	m.nextparkingID++
	m.parking = append(m.parking, p)

	return p
}

func (m *Memoria) ActualizarParqueadero(id string, req modelos.ActualizarParqueaderoRequest) (modelos.Parqueadero, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.parking {
		if p.ID == id {
			if req.Nombre != "" {
				m.parking[i].Nombre = req.Nombre
			}
			if req.Ubicacion != "" {
				m.parking[i].Ubicacion = req.Ubicacion
			}
			if req.Capacidad > 0 {
				m.parking[i].Capacidad = req.Capacidad
			}
			return m.parking[i], true
		}
	}

	return modelos.Parqueadero{}, false
}

func (m *Memoria) EliminarParqueadero(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.parking {
		if p.ID == id {
			m.parking = append(m.parking[:i], m.parking[i+1:]...)
			return true
		}
	}

	return false
}

// -----Espacios-----
func (m *Memoria) ListarEspacios() []modelos.Espacio {
	m.mu.RLock()
	defer m.mu.RUnlock()

	copia := make([]modelos.Espacio, len(m.espacios))
	copy(copia, m.espacios)

	return copia
}

func (m *Memoria) BuscarEspacioPorID(id string) (modelos.Espacio, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, e := range m.espacios {
		if e.ID == id {
			return e, true
		}
	}
	return modelos.Espacio{}, false
}

func (m *Memoria) CrearEspacio(req modelos.CrearEspacioRequest) (modelos.Espacio, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	e := modelos.Espacio{
		ID:            strconv.Itoa(m.nextEspacioID),
		ParqueaderoID: req.ParqueaderoID,
		Numero:        req.Numero,
		Disponible:    true,
	}
	m.nextEspacioID++
	m.espacios = append(m.espacios, e)
	return e, true
}

func (m *Memoria) ActualizarEspacio(id string, req modelos.ActualizarEspacioRequest) (modelos.Espacio, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.espacios {
		if e.ID == id {
			if req.Numero != "" {
				m.espacios[i].Numero = req.Numero
			}
			return m.espacios[i], true
		}
	}
	return modelos.Espacio{}, false
}

func (m *Memoria) EliminarEspacio(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.espacios {
		if e.ID == id {
			m.espacios = append(m.espacios[:i], m.espacios[i+1:]...)
			return true
		}
	}
	return false
}

// -----Ocupaciones-----
func (m *Memoria) ListarOcupaciones() []modelos.Ocupacion {
	m.mu.RLock()
	defer m.mu.RUnlock()
	copia := make([]modelos.Ocupacion, len(m.ocupaciones))
	copy(copia, m.ocupaciones)
	return copia
}

func (m *Memoria) BuscarOcupacionPorID(id string) (modelos.Ocupacion, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, o := range m.ocupaciones {
		if strconv.Itoa(o.ID) == id {
			return o, true
		}
	}
	return modelos.Ocupacion{}, false
}

func (m *Memoria) CrearOcupacion(req modelos.OcuparEspacioRequest) (modelos.Ocupacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	o := modelos.Ocupacion{
		ID:        m.nextOcupacionID,
		EspacioID: req.EspacioID,
		Placa:     req.Placa,
		Entrada:   time.Now(),
		Activa:    true,
	}

	m.nextOcupacionID++
	m.ocupaciones = append(m.ocupaciones, o)

	return o, true
}

func (m *Memoria) ActualizarOcupacion(id string, req modelos.ActualizarOcupacionRequest) (modelos.Ocupacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, o := range m.ocupaciones {
		if strconv.Itoa(o.ID) == id {
			if req.Salida != nil {
				m.ocupaciones[i].Salida = req.Salida
				m.ocupaciones[i].Activa = false
			}
			return m.ocupaciones[i], true
		}
	}
	return modelos.Ocupacion{}, false
}

func (m *Memoria) LiberarOcupacion(id string) (modelos.Ocupacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, o := range m.ocupaciones {
		if strconv.Itoa(o.ID) == id {
			ahora := time.Now()
			m.ocupaciones[i].Salida = &ahora
			m.ocupaciones[i].Activa = false
			return m.ocupaciones[i], true
		}
	}

	return modelos.Ocupacion{}, false
}

// ====================================================================
// ----- MÓDULO: TRANSPORTE INTERNO (CRUD COMPLETO POR ENTIDAD) -----
// ====================================================================

// ----- 1. CRUD: SOLICITUD -----
func (m *Memoria) ListarSolicitudes() []modelos.Solicitud {
	m.mu.RLock()
	defer m.mu.RUnlock()
	copia := make([]modelos.Solicitud, len(m.solicitudes))
	copy(copia, m.solicitudes)
	return copia
}

func (m *Memoria) BuscarSolicitudPorID(id int) (modelos.Solicitud, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, s := range m.solicitudes {
		if s.ID == id {
			return s, true
		}
	}
	return modelos.Solicitud{}, false
}

func (m *Memoria) CrearSolicitud(s modelos.Solicitud) modelos.Solicitud {
	m.mu.Lock()
	defer m.mu.Unlock()
	s.ID = m.nextSolicitudID
	m.nextSolicitudID++
	m.solicitudes = append(m.solicitudes, s)
	return s
}

func (m *Memoria) ActualizarSolicitud(id int, req modelos.Solicitud) (modelos.Solicitud, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.solicitudes {
		if s.ID == id {
			if req.Estado != "" {
				m.solicitudes[i].Estado = req.Estado
			}
			if req.PuntoDestino != "" {
				m.solicitudes[i].PuntoDestino = req.PuntoDestino
			}
			if req.CantPersonas > 0 {
				m.solicitudes[i].CantPersonas = req.CantPersonas
			}
			return m.solicitudes[i], true
		}
	}
	return modelos.Solicitud{}, false
}

func (m *Memoria) EliminarSolicitud(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.solicitudes {
		if s.ID == id {
			m.solicitudes = append(m.solicitudes[:i], m.solicitudes[i+1:]...)
			return true
		}
	}
	return false
}

// ----- 2. CRUD: CARRITO -----
func (m *Memoria) ListarCarritos() []modelos.Carrito {
	m.mu.RLock()
	defer m.mu.RUnlock()
	copia := make([]modelos.Carrito, len(m.carritos))
	copy(copia, m.carritos)
	return copia
}

func (m *Memoria) BuscarCarritoPorID(id int) (modelos.Carrito, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, c := range m.carritos {
		if c.ID == id {
			return c, true
		}
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

func (m *Memoria) ActualizarCarrito(id int, req modelos.Carrito) (modelos.Carrito, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.carritos {
		if c.ID == id {
			if req.NombreCarrito != "" {
				m.carritos[i].NombreCarrito = req.NombreCarrito
			}
			if req.Estado != "" {
				m.carritos[i].Estado = req.Estado
			}
			if req.Capacidad > 0 {
				m.carritos[i].Capacidad = req.Capacidad
			}
			if req.IDRuta > 0 {
				m.carritos[i].IDRuta = req.IDRuta
			}
			return m.carritos[i], true
		}
	}
	return modelos.Carrito{}, false
}

func (m *Memoria) EliminarCarrito(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.carritos {
		if c.ID == id {
			m.carritos = append(m.carritos[:i], m.carritos[i+1:]...)
			return true
		}
	}
	return false
}

// ----- 3. CRUD: LOCACION -----
func (m *Memoria) ListarLocaciones() []modelos.Locacion {
	m.mu.RLock()
	defer m.mu.RUnlock()
	copia := make([]modelos.Locacion, len(m.locaciones))
	copy(copia, m.locaciones)
	return copia
}

func (m *Memoria) BuscarLocacionPorID(id int) (modelos.Locacion, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, l := range m.locaciones {
		if l.ID == id {
			return l, true
		}
	}
	return modelos.Locacion{}, false
}

func (m *Memoria) CrearLocacion(l modelos.Locacion) modelos.Locacion {
	m.mu.Lock()
	defer m.mu.Unlock()
	l.ID = m.nextLocacionID
	l.TimeStamp = time.Now()
	m.nextLocacionID++
	m.locaciones = append(m.locaciones, l)
	return l
}

func (m *Memoria) ActualizarLocacion(id int, req modelos.Locacion) (modelos.Locacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, l := range m.locaciones {
		if l.ID == id {
			if req.Latitud != 0 {
				m.locaciones[i].Latitud = req.Latitud
			}
			if req.Longitud != 0 {
				m.locaciones[i].Longitud = req.Longitud
			}
			m.locaciones[i].TimeStamp = time.Now()
			return m.locaciones[i], true
		}
	}
	return modelos.Locacion{}, false
}

func (m *Memoria) EliminarLocacion(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, l := range m.locaciones {
		if l.ID == id {
			m.locaciones = append(m.locaciones[:i], m.locaciones[i+1:]...)
			return true
		}
	}
	return false
}

// ----- 4. CRUD: RUTA -----
func (m *Memoria) ListarRutas() []modelos.Ruta {
	m.mu.RLock()
	defer m.mu.RUnlock()
	copia := make([]modelos.Ruta, len(m.rutas))
	copy(copia, m.rutas)
	return copia
}

func (m *Memoria) BuscarRutaPorID(id int) (modelos.Ruta, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, r := range m.rutas {
		if r.ID == id {
			return r, true
		}
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

func (m *Memoria) ActualizarRuta(id int, req modelos.Ruta) (modelos.Ruta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, r := range m.rutas {
		if r.ID == id {
			if req.NombreRuta != "" {
				m.rutas[i].NombreRuta = req.NombreRuta
			}
			if req.Descripcion != "" {
				m.rutas[i].Descripcion = req.Descripcion
			}
			return m.rutas[i], true
		}
	}
	return modelos.Ruta{}, false
}

func (m *Memoria) EliminarRuta(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, r := range m.rutas {
		if r.ID == id {
			m.rutas = append(m.rutas[:i], m.rutas[i+1:]...)
			return true
		}
	}
	return false
}

// ----- 5. CRUD: PARADAS -----
func (m *Memoria) ListarParadas() []modelos.Paradas {
	m.mu.RLock()
	defer m.mu.RUnlock()
	copia := make([]modelos.Paradas, len(m.paradas))
	copy(copia, m.paradas)
	return copia
}

func (m *Memoria) BuscarParadaPorID(id int) (modelos.Paradas, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, p := range m.paradas {
		if p.ID == id {
			return p, true
		}
	}
	return modelos.Paradas{}, false
}

func (m *Memoria) CrearParada(p modelos.Paradas) modelos.Paradas {
	m.mu.Lock()
	defer m.mu.Unlock()
	p.ID = m.nextParadaID
	m.nextParadaID++
	m.paradas = append(m.paradas, p)
	return p
}

func (m *Memoria) ActualizarParada(id int, req modelos.Paradas) (modelos.Paradas, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.paradas {
		if p.ID == id {
			if req.Nombre != "" {
				m.paradas[i].Nombre = req.Nombre
			}
			if req.Latitud != 0 {
				m.paradas[i].Latitud = req.Latitud
			}
			if req.Longitud != 0 {
				m.paradas[i].Longitud = req.Longitud
			}
			return m.paradas[i], true
		}
	}
	return modelos.Paradas{}, false
}

	func (m *Memoria) EliminarParada(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.paradas {
		if p.ID == id {
			m.paradas = append(m.paradas[:i], m.paradas[i+1:]...)
			return true
		}
	}
	return false
}