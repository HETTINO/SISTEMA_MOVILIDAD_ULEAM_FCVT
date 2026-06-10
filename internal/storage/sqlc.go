package storage

import (
	"context"
	"database/sql"

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
	return modelos.Espacio{}, false
}

func (a *AlmacenSQLC) CrearEspacio(req modelos.CrearEspacioRequest) (modelos.Espacio, bool) {
	return modelos.Espacio{}, false
}

// Chequeo en tiempo de compilación.
var _ Almacen = (*AlmacenSQLC)(nil)
