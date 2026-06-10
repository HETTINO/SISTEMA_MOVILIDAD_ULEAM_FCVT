package storage

import "SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"

type Almacen interface {
	ListarParqueaderos() []modelos.Parqueadero
	BuscarParqueaderoporID(id string) (modelos.Parqueadero, bool)
	CrearParqueadero(req modelos.CrearParqueaderoRequest) modelos.Parqueadero
	ActualizarParqueadero(id string, req modelos.ActualizarParqueaderoRequest) (modelos.Parqueadero, bool)
	EliminarParqueadero(id string) bool
	// Espacios
	ListarEspacios() []modelos.Espacio
	BuscarEspacioPorID(id string) (modelos.Espacio, bool)
	CrearEspacio(req modelos.CrearEspacioRequest) (modelos.Espacio, bool)
}
