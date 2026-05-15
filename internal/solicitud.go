package internal

// Solicitud: Entidad diseñada por Cristina para el módulo de transporte.
type Solicitud struct {
	ID               uint   `json:"id"`
	CantidadPersonas int    `json:"cantidad_personas"`
	Destino          string `json:"destino"`
	Estado           string `json:"estado"`
}