package storage

import (
	"sync"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"
	"strconv"
)

type Memoria struct {
	parking  []modelos.Parqueadero
	espacios []modelos.Espacio

	nextparkingID int
	nextEspacioID int

	mu sync.RWMutex
}

func NewMemoria() *Memoria {
	return &Memoria{
		parking:       []modelos.Parqueadero{},
		espacios:      []modelos.Espacio{},
		nextEspacioID: 1,
		nextparkingID: 1,
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

// Chequeo en tiempo de compilación: Memoria debe cumplir Almacen.
var _ Almacen = (*Memoria)(nil)
