package modelos

import "time"

// Parqueadero representa un área de estacionamiento dentro de la ULEAM
type Parqueadero struct {
	ID        string `json:"id"`
	Nombre    string `json:"nombre"`
	Ubicacion string `json:"ubicacion"`
	Capacidad int    `json:"capacidad"`
	Activo    bool   `json:"activo"`
}

// Espacio representa un espacio individual dentro de un parqueadero
type Espacio struct {
	ID             string     `json:"id"`
	ParqueaderoID  string     `json:"parqueadero_id"`
	Numero         string     `json:"numero"`
	Disponible     bool       `json:"disponible"`
	ReservadoHasta *time.Time `json:"reservado_hasta,omitempty"`
}

// Ocupacion registra cuando un vehículo ocupa un espacio
type Ocupacion struct {
	ID        string     `json:"id"`
	EspacioID string     `json:"espacio_id"`
	Placa     string     `json:"placa"`
	Entrada   time.Time  `json:"entrada"`
	Salida    *time.Time `json:"salida,omitempty"`
	Activa    bool       `json:"activa"`
}

// --- Request bodies ---

type CrearParqueaderoRequest struct {
	Nombre    string `json:"nombre"`
	Ubicacion string `json:"ubicacion"`
	Capacidad int    `json:"capacidad"`
}

type CrearEspacioRequest struct {
	ParqueaderoID string `json:"parqueadero_id"`
	Numero        string `json:"numero"`
}

type OcuparEspacioRequest struct {
	Placa string `json:"placa"`
}

// --- Response bodies ---

type DisponibilidadResponse struct {
	ParqueaderoID string `json:"parqueadero_id"`
	Nombre        string `json:"nombre"`
	TotalEspacios int    `json:"total_espacios"`
	Disponibles   int    `json:"disponibles"`
	Ocupados      int    `json:"ocupados"`
}
