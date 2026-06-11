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

	// === CAMPOS DEL MÓDULO: ACCESO ===
	usuarios       []modelos.Usuario
	vehiculos      []modelos.Vehiculo
	puntosDeAcceso []modelos.PuntoDeAcceso
	accesos        []modelos.Acceso

	nextPuntoAccesoID int
	nextAccesoID      int
	// ===============================

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

		// === INICIALIZACIÓN DEL MÓDULO ACCESO ===
		usuarios:       []modelos.Usuario{},
		vehiculos:      []modelos.Vehiculo{},
		puntosDeAcceso: []modelos.PuntoDeAcceso{},
		accesos:        []modelos.Acceso{},

		nextPuntoAccesoID: 1,
		nextAccesoID:      1,
		// ===================================
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

// =================================================================
// ----- MÓDULO: ACCESO DE ENTRADA Y SALIDA (MÉTODOS CRUD) -----
// =================================================================

// ListarAccesos devuelve una copia segura de todos los registros de acceso
func (m *Memoria) ListarAccesos() []modelos.Acceso {
	m.mu.RLock()
	defer m.mu.RUnlock()

	copia := make([]modelos.Acceso, len(m.accesos))
	copy(copia, m.accesos)
	return copia
}

// BuscarAccesoPorID busca un registro usando un ID string (convirtiéndolo a int)
func (m *Memoria) BuscarAccesoPorID(id string) (modelos.Acceso, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return modelos.Acceso{}, false
	}

	for _, a := range m.accesos {
		if a.IDAcceso == idInt {
			return a, true
		}
	}
	return modelos.Acceso{}, false
}

// CrearAcceso inserta un nuevo registro de entrada al slice
func (m *Memoria) CrearAcceso(acceso modelos.Acceso) modelos.Acceso {
	m.mu.Lock()
	defer m.mu.Unlock()

	acceso.IDAcceso = m.nextAccesoID
	acceso.TiempoEntrada = time.Now()
	acceso.Estado = "Dentro" // Estado inicial por defecto
	m.nextAccesoID++

	m.accesos = append(m.accesos, acceso)
	return acceso
}

// ActualizarAcceso modifica el estado u observaciones de un acceso existente
func (m *Memoria) ActualizarAcceso(id string, nuevosDatos modelos.Acceso) (modelos.Acceso, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return modelos.Acceso{}, false
	}

	for i, a := range m.accesos {
		if a.IDAcceso == idInt {
			if nuevosDatos.Estado != "" {
				m.accesos[i].Estado = nuevosDatos.Estado
			}
			if nuevosDatos.Observaciones != "" {
				m.accesos[i].Observaciones = nuevosDatos.Observaciones
			}
			if nuevosDatos.TiempoSalida != nil {
				m.accesos[i].TiempoSalida = nuevosDatos.TiempoSalida
			}
			return m.accesos[i], true
		}
	}
	return modelos.Acceso{}, false
}

// EliminarAcceso remueve un registro del histórico utilizando slicing
func (m *Memoria) EliminarAcceso(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	for i, a := range m.accesos {
		if a.IDAcceso == idInt {
			m.accesos = append(m.accesos[:i], m.accesos[i+1:]...)
			return true
		}
	}
	return false
}

// Chequeo en tiempo de compilación: Memoria debe cumplir Almacen.
var _ Almacen = (*Memoria)(nil)
