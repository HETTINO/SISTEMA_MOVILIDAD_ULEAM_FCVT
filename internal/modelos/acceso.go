package modelos

import "time"

// Usuario representa la tabla de usuarios en el sistema
type Usuario struct {
	Cedula     string `json:"cedula"`
	Nombre     string `json:"nombre"`
	Contrasena string `json:"contrasena"`
	Email      string `json:"email"`
	Rol        string `json:"rol"`
}

// Vehiculo representa los datos del vehículo asociado
type Vehiculo struct {
	Placa        string `json:"placa"`
	IDUsuario    string `json:"id_usuario"`
	TipoVehiculo string `json:"tipo_vehiculo"`
	Marca        string `json:"marca"`
	Modelo       string `json:"modelo"`
	Color        string `json:"color"`
	Año          int    `json:"año"`
}

// PuntoDeAcceso representa el lugar físico por donde se ingresa o sale
type PuntoDeAcceso struct {
	IDPuntoAcceso int    `json:"id_punto_acceso"`
	Frecuencia    string `json:"frecuencia"`
	Ubicacion     string `json:"ubicacion"`
}

// Acceso representa el registro de eventos de entrada y salida
type Acceso struct {
	IDAcceso      int        `json:"id_acceso"`
	PlacaVehiculo string     `json:"placa_vehiculo"`
	IDPuntoAcceso int        `json:"id_punto_acceso"`
	TiempoEntrada time.Time  `json:"tiempo_entrada"`
	TiempoSalida  *time.Time `json:"tiempo_salida,omitempty"` // Puntero por si es nulo (aún no sale)
	Estado        string     `json:"estado"`
	Observaciones string     `json:"observaciones"`
}
