package modelos

import "time"

// Solicitud representa la petición de un usuario para usar el transporte interno
type Solicitud struct {
	ID            int    `json:"id_solicitud"`
	CedulaUsuario string `json:"cedula_usuario"`
	IDCarrito     int    `json:"id_carrito"`
	CantPersonas  int    `json:"cant_personas"`
	PuntoDestino  string `json:"punto_destino"`
	Estado        string `json:"estado"`
}

// Carrito representa los vehículos eléctricos/búles de transporte interno
type Carrito struct {
	ID            int    `json:"id_carrito"`
	IDRuta        int    `json:"id_ruta"`
	NombreCarrito string `json:"nombre_carrito"`
	Capacidad     int    `json:"capacidad"`
	Estado        string `json:"estado"`
}

// Locacion representa el rastreo GPS en tiempo real de un carrito
type Locacion struct {
	ID        int       `json:"id_locacion"`
	IDCarrito int       `json:"id_carrito"`
	Latitud   float64   `json:"latitud"`
	Longitud  float64   `json:"longitud"`
	TimeStamp time.Time `json:"time_stamp"`
}

// Ruta representa el camino o circuito establecido en la universidad
type Ruta struct {
	ID          int    `json:"id_ruta"`
	NombreRuta  string `json:"nombre_ruta"`
	Descripcion string `json:"descripcion"`
}

// Paradas representa los puntos fijos donde el carrito se detiene a recoger pasajeros
type Paradas struct {
	ID       int     `json:"id_parada"`
	IDRuta   int     `json:"id_ruta"`
	Nombre   string  `json:"nombre"`
	Latitud  float64 `json:"latitud"`
	Longitud float64 `json:"longitud"`
}
