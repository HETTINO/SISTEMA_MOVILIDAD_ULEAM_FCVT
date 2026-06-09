package storage

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"
)

// ParkingStore almacena los datos en memoria (sin base de datos)
type ParkingStore struct {
	mu           sync.RWMutex
	parqueaderos map[string]modelos.Parqueadero
	espacios     map[string]modelos.Espacio
	ocupaciones  map[string]modelos.Ocupacion
	nextID       int
}

// NewParkingStore crea un store nuevo con datos de ejemplo
func NewParkingStore() *ParkingStore {
	s := &ParkingStore{
		parqueaderos: make(map[string]modelos.Parqueadero),
		espacios:     make(map[string]modelos.Espacio),
		ocupaciones:  make(map[string]modelos.Ocupacion),
		nextID:       1,
	}
	// Semilla: parqueadero de la facultad TI
	p := modelos.Parqueadero{
		ID:        "park-001",
		Nombre:    "Parqueadero FCVT",
		Ubicacion: "Facultad de Ciencias de la Vida y Tecnologías",
		Capacidad: 20,
		Activo:    true,
	}
	s.parqueaderos[p.ID] = p
	// Semilla: 5 espacios
	for i := 1; i <= 5; i++ {
		e := modelos.Espacio{
			ID:            fmt.Sprintf("esp-%03d", i),
			ParqueaderoID: "park-001",
			Numero:        fmt.Sprintf("A%d", i),
			Disponible:    true,
		}
		s.espacios[e.ID] = e
	}
	return s
}

func (s *ParkingStore) generarID(prefix string) string {
	id := fmt.Sprintf("%s-%04d", prefix, s.nextID)
	s.nextID++
	return id
}

// ---------- PARQUEADEROS ----------

func (s *ParkingStore) ListarParqueaderos() []modelos.Parqueadero {
	s.mu.RLock()
	defer s.mu.RUnlock()
	lista := make([]modelos.Parqueadero, 0, len(s.parqueaderos))
	for _, p := range s.parqueaderos {
		lista = append(lista, p)
	}
	return lista
}

func (s *ParkingStore) ObtenerParqueadero(id string) (modelos.Parqueadero, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	p, ok := s.parqueaderos[id]
	if !ok {
		return modelos.Parqueadero{}, errors.New("parqueadero no encontrado")
	}
	return p, nil
}

func (s *ParkingStore) CrearParqueadero(req modelos.CrearParqueaderoRequest) (modelos.Parqueadero, error) {
	if req.Nombre == "" {
		return modelos.Parqueadero{}, errors.New("el campo nombre es requerido")
	}
	if req.Ubicacion == "" {
		return modelos.Parqueadero{}, errors.New("el campo ubicacion es requerido")
	}
	if req.Capacidad <= 0 {
		return modelos.Parqueadero{}, errors.New("la capacidad debe ser mayor a 0")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	p := modelos.Parqueadero{
		ID:        s.generarID("park"),
		Nombre:    req.Nombre,
		Ubicacion: req.Ubicacion,
		Capacidad: req.Capacidad,
		Activo:    true,
	}
	s.parqueaderos[p.ID] = p
	return p, nil
}

func (s *ParkingStore) ActualizarParqueadero(id string, req modelos.CrearParqueaderoRequest) (modelos.Parqueadero, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.parqueaderos[id]
	if !ok {
		return modelos.Parqueadero{}, errors.New("parqueadero no encontrado")
	}
	if req.Nombre != "" {
		p.Nombre = req.Nombre
	}
	if req.Ubicacion != "" {
		p.Ubicacion = req.Ubicacion
	}
	if req.Capacidad > 0 {
		p.Capacidad = req.Capacidad
	}
	s.parqueaderos[id] = p
	return p, nil
}

func (s *ParkingStore) EliminarParqueadero(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.parqueaderos[id]; !ok {
		return errors.New("parqueadero no encontrado")
	}
	delete(s.parqueaderos, id)
	return nil
}

// ---------- ESPACIOS ----------

func (s *ParkingStore) ListarEspacios(parqueaderoID string) []modelos.Espacio {
	s.mu.RLock()
	defer s.mu.RUnlock()
	lista := make([]modelos.Espacio, 0)
	for _, e := range s.espacios {
		if parqueaderoID == "" || e.ParqueaderoID == parqueaderoID {
			lista = append(lista, e)
		}
	}
	return lista
}

func (s *ParkingStore) ObtenerEspacio(id string) (modelos.Espacio, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, ok := s.espacios[id]
	if !ok {
		return modelos.Espacio{}, errors.New("espacio no encontrado")
	}
	return e, nil
}

func (s *ParkingStore) CrearEspacio(req modelos.CrearEspacioRequest) (modelos.Espacio, error) {
	if req.ParqueaderoID == "" {
		return modelos.Espacio{}, errors.New("el campo parqueadero_id es requerido")
	}
	if req.Numero == "" {
		return modelos.Espacio{}, errors.New("el campo numero es requerido")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.parqueaderos[req.ParqueaderoID]; !ok {
		return modelos.Espacio{}, errors.New("parqueadero no encontrado")
	}
	e := modelos.Espacio{
		ID:            s.generarID("esp"),
		ParqueaderoID: req.ParqueaderoID,
		Numero:        req.Numero,
		Disponible:    true,
	}
	s.espacios[e.ID] = e
	return e, nil
}

func (s *ParkingStore) ReservarEspacio(id string) (modelos.Espacio, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.espacios[id]
	if !ok {
		return modelos.Espacio{}, errors.New("espacio no encontrado")
	}
	if !e.Disponible {
		return modelos.Espacio{}, errors.New("el espacio ya está ocupado")
	}
	hasta := time.Now().Add(5 * time.Minute)
	e.Disponible = false
	e.ReservadoHasta = &hasta
	s.espacios[id] = e
	return e, nil
}

func (s *ParkingStore) LiberarEspacio(id string) (modelos.Espacio, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.espacios[id]
	if !ok {
		return modelos.Espacio{}, errors.New("espacio no encontrado")
	}
	e.Disponible = true
	e.ReservadoHasta = nil
	s.espacios[id] = e
	return e, nil
}

func (s *ParkingStore) EliminarEspacio(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.espacios[id]; !ok {
		return errors.New("espacio no encontrado")
	}
	delete(s.espacios, id)
	return nil
}

// ---------- DISPONIBILIDAD ----------

func (s *ParkingStore) ObtenerDisponibilidad(parqueaderoID string) (modelos.DisponibilidadResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	p, ok := s.parqueaderos[parqueaderoID]
	if !ok {
		return modelos.DisponibilidadResponse{}, errors.New("parqueadero no encontrado")
	}
	total, disponibles := 0, 0
	for _, e := range s.espacios {
		if e.ParqueaderoID == parqueaderoID {
			total++
			if e.Disponible {
				disponibles++
			}
		}
	}
	return modelos.DisponibilidadResponse{
		ParqueaderoID: p.ID,
		Nombre:        p.Nombre,
		TotalEspacios: total,
		Disponibles:   disponibles,
		Ocupados:      total - disponibles,
	}, nil
}

// ---------- OCUPACIONES ----------

func (s *ParkingStore) ListarOcupaciones() []modelos.Ocupacion {
	s.mu.RLock()
	defer s.mu.RUnlock()
	lista := make([]modelos.Ocupacion, 0, len(s.ocupaciones))
	for _, o := range s.ocupaciones {
		lista = append(lista, o)
	}
	return lista
}

func (s *ParkingStore) RegistrarOcupacion(req modelos.OcuparEspacioRequest, espacioID string) (modelos.Ocupacion, error) {
	if req.Placa == "" {
		return modelos.Ocupacion{}, errors.New("el campo placa es requerido")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.espacios[espacioID]
	if !ok {
		return modelos.Ocupacion{}, errors.New("espacio no encontrado")
	}
	if !e.Disponible {
		return modelos.Ocupacion{}, errors.New("el espacio no está disponible")
	}
	e.Disponible = false
	s.espacios[espacioID] = e
	o := modelos.Ocupacion{
		ID:        s.generarID("ocu"),
		EspacioID: espacioID,
		Placa:     req.Placa,
		Entrada:   time.Now(),
		Activa:    true,
	}
	s.ocupaciones[o.ID] = o
	return o, nil
}
