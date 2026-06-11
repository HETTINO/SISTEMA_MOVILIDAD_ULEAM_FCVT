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
	ActualizarEspacio(id string, req modelos.ActualizarEspacioRequest) (modelos.Espacio, bool)
	EliminarEspacio(id string) bool

	// Ocupaciones
	ListarOcupaciones() []modelos.Ocupacion
	BuscarOcupacionPorID(id string) (modelos.Ocupacion, bool)
	CrearOcupacion(req modelos.OcuparEspacioRequest) (modelos.Ocupacion, bool)
	LiberarOcupacion(id string) (modelos.Ocupacion, bool)

// =========================================================
    // SOLICITUDES DE TRANSPORTE - CRISTINA
    // =========================================================
    ListarSolicitudes() []modelos.Solicitud
    BuscarSolicitudPorID(id string) (modelos.Solicitud, bool)
    CrearSolicitud(s modelos.Solicitud) modelos.Solicitud
    ActualizarSolicitud(id string, datos modelos.Solicitud) (modelos.Solicitud, bool)
    BorrarSolicitud(id string) bool

    // =========================================================
    // RUTAS
    // =========================================================
    ListarRutas() []modelos.Ruta
    BuscarRutaPorID(id string) (modelos.Ruta, bool)
    CrearRuta(r modelos.Ruta) modelos.Ruta
    ActualizarRuta(id string, datos modelos.Ruta) (modelos.Ruta, bool)
    BorrarRuta(id string) bool

    // =========================================================
    // PARADAS
    // =========================================================
    ListarParadas() []modelos.Parada
    BuscarParadaPorID(id string) (modelos.Parada, bool)
    CrearParada(p modelos.Parada) modelos.Parada
    ActualizarParada(id string, datos modelos.Parada) (modelos.Parada, bool)
    BorrarParada(id string) bool

    // =========================================================
    // CARRITOS
    // =========================================================
    ListarCarritos() []modelos.Carrito
    BuscarCarritoPorID(id string) (modelos.Carrito, bool)
    CrearCarrito(c modelos.Carrito) modelos.Carrito
    ActualizarCarrito(id string, datos modelos.Carrito) (modelos.Carrito, bool)
    BorrarCarrito(id string) bool

    // =========================================================
    // LOCACIÓN
    // =========================================================
    ListarLocaciones() []modelos.Locacion
    BuscarLocacionPorID(id string) (modelos.Locacion, bool)
    CrearLocacion(l modelos.Locacion) modelos.Locacion
    ActualizarLocacion(id string, datos modelos.Locacion) (modelos.Locacion, bool)
    BorrarLocacion(id string) bool
}

