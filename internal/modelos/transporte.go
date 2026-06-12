package modelos

import "time"

// Conductor representa al chofer asignado a una unidad de transporte
type Conductor struct {
	CedulaConductor string `json:"cedula_conductor"`
	Nombre          string `json:"nombre"`
	Licencia        string `json:"licencia"` // Tipo de licencia (ej: Profesional Tipo E)
	Telefono        string `json:"telefono"`
	Estado          string `json:"estado"`   // Activo, En Ruta, Libre
}

// Ruta representa el trayecto predefinido del transporte institucional
type Ruta struct {
	IDRuta      int    `json:"id_ruta"`
	NombreRuta  string `json:"nombre_ruta"`  // Ej: "Ruta Manta - Portoviejo" o "Interna ULEAM"
	Origen      string `json:"origen"`
	Destino     string `json:"destino"`
	Precio      float64 `json:"precio"`       // Por si tiene costo o es gratuito (0.00)
	Estado      string `json:"estado"`       // Activa, Suspendida
}

// UnidadTransporte representa el bus o vehículo físico de la institución
type UnidadTransporte struct {
	PlacaUnidad     string `json:"placa_unidad"`
	Disco           string `json:"disco"`            // Número de unidad (ej: "Bus 05")
	Capacidad       int    `json:"capacidad"`        // Cantidad de pasajeros sentados
	CedulaConductor string `json:"cedula_conductor"` // Chofer asignado fijo
	Marca           string `json:"marca"`
	Modelo          string `json:"modelo"`
	Estado          string `json:"estado"`           // Operativo, En Mantenimiento
}

// Viaje representa la tabla transaccional (los recorridos en tiempo real)
type Viaje struct {
	IDViaje         int        `json:"id_viaje"`
	IDRuta          int        `json:"id_ruta"`
	PlacaUnidad     string     `json:"placa_unidad"`
	CedulaConductor string     `json:"cedula_conductor"`
	HoraSalidaPlan  time.Time  `json:"hora_salida_planificado"`
	HoraLlegadaReal *time.Time `json:"hora_llegada_real,omitempty"` // Puntero por si está en camino (nulo)
	PasajerosA bordo int        `json:"pasajeros_a_bordo"`
	Estado          string     `json:"estado"`                      // Programado, En Ruta, Finalizado, Cancelado
}