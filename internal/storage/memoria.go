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

// ============================================================================
// MODULO: TRANSPORTE INTERNO - MÉTODOS SEED INDIVIDUALES
// ============================================================================

// SeedRutas carga los circuitos de transporte iniciales en memoria
func (m *Memoria) SeedRutas() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.rutas = []modelos.Ruta{
		{ID: 1, Nombre: "Ruta Troncal Perimetral", Descripcion: "Circuito externo desde la Entrada Principal hasta Ciencias Médicas."},
		{ID: 2, Nombre: "Ruta Tecnológica FCVT", Descripcion: "Conecta el Bloque de Aulas FCVT con los Laboratorios de Cómputo."},
		{ID: 3, Nombre: "Ruta Administrativa", Descripcion: "Recorrido express entre Rectorado, Biblioteca Central y el Auditorio."},
	}
	m.nextRutaID = 4
}

// SeedParadas carga los puntos de control en las rutas (Atención a 'Longitu' sin 'd')
func (m *Memoria) SeedParadas() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.paradas = []modelos.Parada{
		{IDParada: 1, Nombre: "Parada Entrada Principal", Latitud: -0.9510, Longitu: -80.7020, RutaID: 1},
		{IDParada: 2, Nombre: "Parada Facultad Ciencias Médicas", Latitud: -0.9565, Longitu: -80.7095, RutaID: 1},
		{IDParada: 3, Nombre: "Parada Bloque Aulas FCVT", Latitud: -0.9525, Longitu: -80.7035, RutaID: 2},
		{IDParada: 4, Nombre: "Parada Laboratorios de Cómputo", Latitud: -0.9530, Longitu: -80.7040, RutaID: 2},
		{IDParada: 5, Nombre: "Parada Rectorado", Latitud: -0.9505, Longitu: -80.7012, RutaID: 3},
	}
	m.nextParadaID = 6
}

// SeedCarritos carga la flota de vehículos eléctricos simulados
func (m *Memoria) SeedCarritos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.carritos = []modelos.Carrito{
		{ID: 1, NombreCarrito: "Eco-Carrito FCVT #1", Capacidad: 5, Estado: "Disponible", RutaID: 2},
		{ID: 2, NombreCarrito: "Eco-Carrito Manta #2", Capacidad: 5, Estado: "En Ruta", RutaID: 1},
		{ID: 3, NombreCarrito: "Eco-Carrito Central #3", Capacidad: 4, Estado: "Mantenimiento", RutaID: 3},
	}
	m.nextCarritoID = 4
}

// SeedLocaciones carga las coordenadas iniciales de geolocalización de la flota
func (m *Memoria) SeedLocaciones() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.locaciones = []modelos.Locacion{
		{ID: 1, Latitud: -0.9525, Longitud: -80.7035, TimeStamp: time.Now(), CarritoID: "1"},
		{ID: 2, Latitud: -0.9510, Longitud: -80.7020, TimeStamp: time.Now().Add(-5 * time.Minute), CarritoID: "2"},
	}
	m.nextLocacionID = 3
}

// SeedSolicitudes carga las peticiones de transporte que revisarás en Postman
func (m *Memoria) SeedSolicitudes() {
	m.mu.Lock()
	defer m.mu.Unlock()

	idCarro1 := "1"
	idCarro2 := "2"

	m.solicitudes = []modelos.Solicitud{
		{
			ID:            1,
			CedulaUsuario: "1312345678",
			CantPersonas:  2,
			PuntoDestino:  "Bloque de Aulas FCVT",
			Estado:        "Pendiente",
			IDCarrito:     nil,
		},
		{
			ID:            2,
			CedulaUsuario: "1315432109",
			CantPersonas:  1,
			PuntoDestino:  "Laboratorios de Simulación",
			Estado:        "Asignado",
			IDCarrito:     &idCarro1,
		},
		{
			ID:            3,
			CedulaUsuario: "1723456781",
			CantPersonas:  4,
			PuntoDestino:  "Facultad de Ciencias Médicas",
			Estado:        "En Camino",
			IDCarrito:     &idCarro2,
		},
		{
			ID:            4,
			CedulaUsuario: "1309876543",
			CantPersonas:  3,
			PuntoDestino:  "Auditorio Principal",
			Estado:        "Finalizado",
			IDCarrito:     &idCarro1,
		},
	}
	m.nextSolicitudID = 5
}

// BuscarSolicitudPorID localiza una solicitud y devuelve (modelos.Solicitud, bool)
func (m *Memoria) BuscarSolicitudPorID(id string) (modelos.Solicitud, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, s := range m.solicitudes {
		if strconv.Itoa(s.ID) == id {
			return s, true
		}
	}
	return modelos.Solicitud{}, false
}

// CrearSolicitud recibe un Request, genera el ID autoincremental y guarda el registro
func (m *Memoria) CrearSolicitud(req modelos.Solicitud) (modelos.Solicitud, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	s := modelos.Solicitud{
		ID:            m.nextSolicitudID,
		CedulaUsuario: req.CedulaUsuario,
		CantPersonas:  req.CantPersonas,
		PuntoDestino:  req.PuntoDestino,
		Estado:        "Pendiente", // Estado inicial por defecto
		IDCarrito:     nil,         // Inicialmente sin carrito asignado
	}

	m.nextSolicitudID++
	m.solicitudes = append(m.solicitudes, s)

	return s, true
}

// ActualizarSolicitud busca por ID, modifica los campos si no vienen vacíos y retorna un bool
func (m *Memoria) ActualizarSolicitud(id string, req modelos.Solicitud) (modelos.Solicitud, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.solicitudes {
		if strconv.Itoa(s.ID) == id {
			if req.CedulaUsuario != "" {
				m.solicitudes[i].CedulaUsuario = req.CedulaUsuario
			}
			if req.CantPersonas > 0 {
				m.solicitudes[i].CantPersonas = req.CantPersonas
			}
			if req.PuntoDestino != "" {
				m.solicitudes[i].PuntoDestino = req.PuntoDestino
			}
			if req.Estado != "" {
				m.solicitudes[i].Estado = req.Estado
			}
			// Permite actualizar el puntero del carrito asignado
			m.solicitudes[i].IDCarrito = req.IDCarrito

			return m.solicitudes[i], true
		}
	}
	return modelos.Solicitud{}, false
}

// EliminarSolicitud remueve la solicitud del slice usando el patrón de recorte append
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