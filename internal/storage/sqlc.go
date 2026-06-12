package storage

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/google/uuid"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/storage/sqlcdb"
)

type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

func NewAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{q: sqlcdb.New(db)}
}

func aParqueaderoSQLC(p sqlcdb.Parqueadero) modelos.Parqueadero {
	return modelos.Parqueadero{
		ID:        p.ID,
		Nombre:    p.Nombre,
		Ubicacion: p.Ubicacion,
		Capacidad: int(p.Capacidad),
		Activo:    p.Activo,
	}
}

func (a *AlmacenSQLC) ListarParqueaderos() []modelos.Parqueadero {
	filas, err := a.q.ListarParqueaderos(context.Background())
	if err != nil {
		return []modelos.Parqueadero{}
	}
	out := make([]modelos.Parqueadero, 0, len(filas))
	for _, f := range filas {
		out = append(out, aParqueaderoSQLC(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarParqueaderoporID(id string) (modelos.Parqueadero, bool) {
	fila, err := a.q.ObtenerParqueadero(context.Background(), id)
	if err != nil {
		return modelos.Parqueadero{}, false
	}
	return aParqueaderoSQLC(fila), true
}

func (a *AlmacenSQLC) CrearParqueadero(req modelos.CrearParqueaderoRequest) modelos.Parqueadero {
	err := a.q.CrearParqueadero(context.Background(), sqlcdb.CrearParqueaderoParams{
		ID:        "",
		Nombre:    req.Nombre,
		Ubicacion: req.Ubicacion,
		Capacidad: int64(req.Capacidad),
		Activo:    req.Activo,
	})
	if err != nil {
		return modelos.Parqueadero{}
	}
	return modelos.Parqueadero{
		Nombre:    req.Nombre,
		Ubicacion: req.Ubicacion,
		Capacidad: req.Capacidad,
		Activo:    req.Activo,
	}
}

func (a *AlmacenSQLC) ActualizarParqueadero(id string, req modelos.ActualizarParqueaderoRequest) (modelos.Parqueadero, bool) {
	// No hay query generada para actualizar; retornamos false por ahora
	return modelos.Parqueadero{}, false
}

func (a *AlmacenSQLC) EliminarParqueadero(id string) bool {
	err := a.q.EliminarParqueadero(context.Background(), id)
	return err == nil
}

func (a *AlmacenSQLC) ListarEspacios() []modelos.Espacio {
	filas, err := a.q.ListarEspacios(context.Background())
	if err != nil {
		return []modelos.Espacio{}
	}
	out := make([]modelos.Espacio, 0, len(filas))
	for _, f := range filas {
		out = append(out, modelos.Espacio{
			ID:            f.ID,
			ParqueaderoID: f.ParqueaderoID,
			Numero:        f.Numero,
			Disponible:    f.Disponible,
		})
	}
	return out
}

func (a *AlmacenSQLC) BuscarEspacioPorID(id string) (modelos.Espacio, bool) {

	e, err := a.q.ObtenerEspacio(context.Background(), id)

	if err != nil {
		return modelos.Espacio{}, false
	}

	return modelos.Espacio{
		ID:            e.ID,
		ParqueaderoID: e.ParqueaderoID,
		Numero:        e.Numero,
		Disponible:    e.Disponible,
	}, true
}
func (a *AlmacenSQLC) CrearEspacio(req modelos.CrearEspacioRequest) (modelos.Espacio, bool) {

	err := a.q.CrearEspacio(context.Background(), sqlcdb.CrearEspacioParams{
		ID:             uuid.New().String(),
		ParqueaderoID:  req.ParqueaderoID,
		Numero:         req.Numero,
		Disponible:     true,
		ReservadoHasta: sql.NullTime{},
	})

	if err != nil {
		return modelos.Espacio{}, false
	}

	return modelos.Espacio{
		ParqueaderoID: req.ParqueaderoID,
		Numero:        req.Numero,
		Disponible:    true,
	}, true
}
func (a *AlmacenSQLC) ActualizarEspacio(id string, req modelos.ActualizarEspacioRequest) (modelos.Espacio, bool) {

	e, err := a.q.ActualizarEspacio(context.Background(),
		sqlcdb.ActualizarEspacioParams{
			ID:         id,
			Numero:     req.Numero,
			Disponible: req.Disponible,
		})

	if err != nil {
		return modelos.Espacio{}, false
	}

	return modelos.Espacio{
		ID:            e.ID,
		ParqueaderoID: e.ParqueaderoID,
		Numero:        e.Numero,
		Disponible:    e.Disponible,
	}, true
}

func (a *AlmacenSQLC) EliminarEspacio(id string) bool {
	return a.q.EliminarEspacio(context.Background(), id) == nil
}

//Ocupaciones

func (a *AlmacenSQLC) ListarOcupaciones() []modelos.Ocupacion {

	filas, err := a.q.ListarOcupaciones(context.Background())

	if err != nil {
		return []modelos.Ocupacion{}
	}

	var out []modelos.Ocupacion

	for _, f := range filas {

		var salida *time.Time

		if f.Salida.Valid {
			salida = &f.Salida.Time
		}

		out = append(out, modelos.Ocupacion{
			ID:        atoiSeguro(f.ID),
			EspacioID: f.EspacioID,
			Placa:     f.Placa,
			Entrada:   f.Entrada,
			Salida:    salida,
			Activa:    f.Activa,
		})
	}

	return out
}

func (a *AlmacenSQLC) BuscarOcupacionPorID(id string) (modelos.Ocupacion, bool) {

	f, err := a.q.ObtenerOcupacion(context.Background(), id)

	if err != nil {
		return modelos.Ocupacion{}, false
	}

	var salida *time.Time

	if f.Salida.Valid {
		salida = &f.Salida.Time
	}

	return modelos.Ocupacion{
		ID:        atoiSeguro(f.ID),
		EspacioID: f.EspacioID,
		Placa:     f.Placa,
		Entrada:   f.Entrada,
		Salida:    salida,
		Activa:    f.Activa,
	}, true
}

func (a *AlmacenSQLC) CrearOcupacion(req modelos.OcuparEspacioRequest) (modelos.Ocupacion, bool) {

	id := uuid.New().String()

	err := a.q.CrearOcupacion(context.Background(),
		sqlcdb.CrearOcupacionParams{
			ID:        id,
			EspacioID: req.EspacioID,
			Placa:     req.Placa,
			Entrada:   time.Now(),
			Salida:    sql.NullTime{},
			Activa:    true,
		})

	if err != nil {
		return modelos.Ocupacion{}, false
	}

	return modelos.Ocupacion{
		EspacioID: req.EspacioID,
		Placa:     req.Placa,
		Entrada:   time.Now(),
		Activa:    true,
	}, true
}

func (a *AlmacenSQLC) LiberarOcupacion(id string) (modelos.Ocupacion, bool) {

	ahora := time.Now()

	f, err := a.q.ActualizarOcupacion(context.Background(),
		sqlcdb.ActualizarOcupacionParams{
			ID: id,
			Salida: sql.NullTime{
				Time:  ahora,
				Valid: true,
			},
			Activa: false,
		})

	if err != nil {
		return modelos.Ocupacion{}, false
	}

	return modelos.Ocupacion{
		ID:        atoiSeguro(f.ID),
		EspacioID: f.EspacioID,
		Placa:     f.Placa,
		Entrada:   f.Entrada,
		Salida:    &ahora,
		Activa:    false,
	}, true
}

func atoiSeguro(valor string) int {
	n, _ := strconv.Atoi(valor)
	return n
}


// ============================================================================
// SOLICITUDES DE TRANSPORTE INTERNO (MELANIE)
// ============================================================================

func (a *AlmacenSQLite) ListarSolicitudes() []modelos.Solicitud {
	var solicitudes []modelos.Solicitud
	a.db.Find(&solicitudes)
	return solicitudes
}

func (a *AlmacenSQLite) BuscarSolicitudPorID(id string) (modelos.Solicitud, bool) {
	var s modelos.Solicitud
	if err := a.db.First(&s, "id = ?", id).Error; err != nil {
		return modelos.Solicitud{}, false
	}
	return s, true
}

func (a *AlmacenSQLite) CrearSolicitud(s modelos.Solicitud) modelos.Solicitud {
	a.db.Create(&s)
	return s
}

func (a *AlmacenSQLite) ActualizarSolicitud(id string, datos modelos.Solicitud) (modelos.Solicitud, bool) {
	var s modelos.Solicitud
	if err := a.db.First(&s, "id = ?", id).Error; err != nil {
		return modelos.Solicitud{}, false
	}
	
	// Utiliza el método Updates de GORM tal como lo hace tu equipo
	a.db.Model(&s).Updates(datos)
	return s, true
}

func (a *AlmacenSQLite) BorrarSolicitud(id string) bool {
	res := a.db.Where("id = ?", id).Delete(&modelos.Solicitud{})
	return res.RowsAffected > 0
}

// Chequeo en tiempo de compilación.
var _ Almacen = (*AlmacenSQLC)(nil)
