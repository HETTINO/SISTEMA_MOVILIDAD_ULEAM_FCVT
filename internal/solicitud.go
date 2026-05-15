package internal

// Solicitud: Entidad diseñada por Cristina para el módulo de transporte.

type Solicitud struct {
	ID             uint   `json:"id_solicitud"`
	CedulaUsuario  string `json:"fk_cedula_usuario"`
	IDCarrito      int    `json:"fk_id_carrito"`
	CantPersonas   int    `json:"cant_personas"`
	PuntoDestino   string `json:"punto_destino"`
	Estado         string `json:"estado"`
}