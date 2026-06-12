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
	mu              sync.RWMutex


	// MODULO: TRANSPORTE INTERNO CRISTINA
	rutas             []modelos.Ruta
	paradas           []modelos.Parada
	solicitudes       []modelos.Solicitud
	carritos          []modelos.Carrito
	locaciones        []modelos.Locacion

	nextRutaID        int
	nextSolicitudID   int
	nextParadaID      int
	nextCarritoID     int
	nextLocacionID    int
}



func NewMemoria() *Memoria {
	return &Memoria{
		parking:       []modelos.Parqueadero{},
		espacios:      []modelos.Espacio{},
		nextEspacioID: 1,
		nextparkingID: 1,

		// MODULO: TRANSPORTE INTERNO CRISTINA
		rutas:           []modelos.Ruta{},
		paradas:         []modelos.Parada{},
		solicitudes:     []modelos.Solicitud{},
		carritos:        []modelos.Carrito{},
		locaciones:      []modelos.Locacion{},

		nextRutaID:      1,
		nextSolicitudID: 1,
		nextParadaID:    1,
		nextCarritoID:   1,
		nextLocacionID:  1,
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
	m.nextOcupacionID = 3
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

// ListarEspacios, BuscarEspacioPorID, CrearEspacio
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


// =========================================================
// MODULO: TRANSPORTE INTERNO CRISTINA (5 ENTIDADES AUTOINCREMENTALES)
// =========================================================

// ----- 1. RUTAS -----
func (m *Memoria) ListarRutas() []modelos.Ruta {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]modelos.Ruta, len(m.rutas))
	copy(copia, m.rutas)
	return copia
}

func (m *Memoria) BuscarRutaPorID(id string) (modelos.Ruta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, r := range m.rutas {
		if strconv.Itoa(r.ID) == id {
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

func (m *Memoria) ActualizarRuta(id string, datos modelos.Ruta) (modelos.Ruta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, r := range m.rutas {
		if strconv.Itoa(r.ID) == id {
			datos.ID = r.ID // Mantiene su ID numérico original
			m.rutas[i] = datos
			return datos, true
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

// ----- 2. PARADAS -----
func (m *Memoria) ListarParadas() []modelos.Parada {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]modelos.Parada, len(m.paradas))
	copy(copia, m.paradas)
	return copia
}

func (m *Memoria) BuscarParadaPorID(id string) (modelos.Parada, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, p := range m.paradas {
		if strconv.Itoa(p.IDParada) == id {
			return p, true
		}
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
			datos.IDParada = p.IDParada
			m.paradas[i] = datos
			return datos, true
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

// ----- 3. CARRITOS -----
func (m *Memoria) ListarCarritos() []modelos.Carrito {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]modelos.Carrito, len(m.carritos))
	copy(copia, m.carritos)
	return copia
}

func (m *Memoria) BuscarCarritoPorID(id string) (modelos.Carrito, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, c := range m.carritos {
		if strconv.Itoa(c.ID) == id {
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

func (m *Memoria) ActualizarCarrito(id string, datos modelos.Carrito) (modelos.Carrito, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.carritos {
		if strconv.Itoa(c.ID) == id {
			datos.ID = c.ID
			m.carritos[i] = datos
			return datos, true
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

// ----- 4. LOCACIONES -----
func (m *Memoria) ListarLocaciones() []modelos.Locacion {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]modelos.Locacion, len(m.locaciones))
	copy(copia, m.locaciones)
	return copia
}

func (m *Memoria) BuscarLocacionPorID(id string) (modelos.Locacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.locaciones {
		if strconv.Itoa(l.ID) == id {
			return l, true
		}
	}
	return modelos.Locacion{}, false
}

func (m *Memoria) CrearLocacion(l modelos.Locacion) modelos.Locacion {
	m.mu.Lock()
	defer m.mu.Unlock()
	l.ID = m.nextLocacionID
	m.nextLocacionID++
	m.locaciones = append(m.locaciones, l)
	return l
}

func (m *Memoria) ActualizarLocacion(id string, datos modelos.Locacion) (modelos.Locacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, l := range m.locaciones {
		if strconv.Itoa(l.ID) == id {
			datos.ID = l.ID
			m.locaciones[i] = datos
			return datos, true
		}
	}
	return modelos.Locacion{}, false
}

func (m *Memoria) BorrarLocacion(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, l := range m.locaciones {
		if strconv.Itoa(l.ID) == id {
			m.locaciones = append(m.locaciones[:i], m.locaciones[i+1:]...)
			return true
		}
	}
	return false
}

// ----- 5. SOLICITUDES -----
func (m *Memoria) ListarSolicitudes() []modelos.Solicitud {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]modelos.Solicitud, len(m.solicitudes))
	copy(copia, m.solicitudes)
	return copia
}

func (m *Memoria) BuscarSolicitudPorID(id string) (modelos.Solicitud, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, s := range m.solicitudes {
		if strconv.Itoa(s.ID) == id {
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

func (m *Memoria) ActualizarSolicitud(id string, datos modelos.Solicitud) (modelos.Solicitud, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.solicitudes {
		if strconv.Itoa(s.ID) == id {
			datos.ID = s.ID
			m.solicitudes[i] = datos
			return datos, true
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


// Chequeo en tiempo de compilación: Memoria debe cumplir Almacen.
var _ Almacen = (*Memoria)(nil)
