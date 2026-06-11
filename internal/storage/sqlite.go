package storage

import (
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// AlmacenSQLite implementa la interfaz Almacen usando GORM sobre SQLite.
type AlmacenSQLite struct {
	db *gorm.DB
}

// NuevoAlmacenSQLite envuelve una conexión *gorm.DB ya abierta.
func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

// =========================================================
// PARQUEADEROS
// =========================================================

func (a *AlmacenSQLite) ListarParqueaderos() []modelos.Parqueadero {
	var parqueaderos []modelos.Parqueadero
	a.db.Find(&parqueaderos)
	return parqueaderos
}

func (a *AlmacenSQLite) BuscarParqueaderoporID(id string) (modelos.Parqueadero, bool) {
	var p modelos.Parqueadero
	if err := a.db.First(&p, "id = ?", id).Error; err != nil {
		return modelos.Parqueadero{}, false
	}
	return p, true
}

func (a *AlmacenSQLite) CrearParqueadero(req modelos.CrearParqueaderoRequest) modelos.Parqueadero {
	var ultimo modelos.Parqueadero
	a.db.Order("id DESC").First(&ultimo)

	nuevoID := 1
	if ultimo.ID != "" {
		fmt.Sscanf(ultimo.ID, "%d", &nuevoID)
		nuevoID++
	}

	p := modelos.Parqueadero{
		ID:        fmt.Sprintf("%d", nuevoID),
		Nombre:    req.Nombre,
		Ubicacion: req.Ubicacion,
		Capacidad: req.Capacidad,
		Activo:    true,
	}
	a.db.Create(&p)
	return p
}

func (a *AlmacenSQLite) ActualizarParqueadero(id string, req modelos.ActualizarParqueaderoRequest) (modelos.Parqueadero, bool) {
	var existente modelos.Parqueadero
	if err := a.db.First(&existente, "id = ?", id).Error; err != nil {
		return modelos.Parqueadero{}, false
	}
	existente.Nombre = req.Nombre
	existente.Ubicacion = req.Ubicacion
	existente.Capacidad = req.Capacidad
	existente.Activo = req.Activo
	a.db.Save(&existente)
	return existente, true
}

func (a *AlmacenSQLite) EliminarParqueadero(id string) bool {
	res := a.db.Where("id = ?", id).Delete(&modelos.Parqueadero{})
	return res.RowsAffected > 0
}

// =========================================================
// ESPACIOS
// =========================================================

func (a *AlmacenSQLite) ListarEspacios() []modelos.Espacio {
	var espacios []modelos.Espacio
	a.db.Find(&espacios)
	return espacios
}

func (a *AlmacenSQLite) BuscarEspacioPorID(id string) (modelos.Espacio, bool) {
	var e modelos.Espacio
	if err := a.db.First(&e, "id = ?", id).Error; err != nil {
		return modelos.Espacio{}, false
	}
	return e, true
}

func (a *AlmacenSQLite) CrearEspacio(req modelos.CrearEspacioRequest) (modelos.Espacio, bool) {
	e := modelos.Espacio{
		ID:            req.ID, // ← usar el ID del request
		ParqueaderoID: req.ParqueaderoID,
		Numero:        req.Numero,
		Disponible:    true,
	}
	res := a.db.Create(&e)
	return e, res.Error == nil
}

func (a *AlmacenSQLite) ActualizarEspacio(id string, req modelos.ActualizarEspacioRequest) (modelos.Espacio, bool) {
	var e modelos.Espacio

	if err := a.db.First(&e, "id = ?", id).Error; err != nil {
		return modelos.Espacio{}, false
	}

	e.Numero = req.Numero
	e.Disponible = req.Disponible

	a.db.Save(&e)

	return e, true
}

func (a *AlmacenSQLite) EliminarEspacio(id string) bool {
	res := a.db.Where("id = ?", id).Delete(&modelos.Espacio{})
	return res.RowsAffected > 0
}

func (a *AlmacenSQLite) ListarOcupaciones() []modelos.Ocupacion {
	var ocupaciones []modelos.Ocupacion
	a.db.Find(&ocupaciones)
	return ocupaciones
}

func (a *AlmacenSQLite) BuscarOcupacionPorID(id string) (modelos.Ocupacion, bool) {
	var o modelos.Ocupacion

	if err := a.db.First(&o, "id = ?", id).Error; err != nil {
		return modelos.Ocupacion{}, false
	}

	return o, true
}

func (a *AlmacenSQLite) CrearOcupacion(req modelos.OcuparEspacioRequest) (modelos.Ocupacion, bool) {

	var ultima modelos.Ocupacion
	a.db.Order("id DESC").First(&ultima)

	o := modelos.Ocupacion{
		ID:        req.ID, // ← usar el del request
		EspacioID: req.EspacioID,
		Placa:     req.Placa,
		Entrada:   time.Now(),
		Activa:    true,
	}

	res := a.db.Create(&o)

	return o, res.Error == nil
}

func (a *AlmacenSQLite) LiberarOcupacion(id string) (modelos.Ocupacion, bool) {

	var o modelos.Ocupacion

	if err := a.db.First(&o, "id = ?", id).Error; err != nil {
		return modelos.Ocupacion{}, false
	}

	ahora := time.Now()

	o.Salida = &ahora
	o.Activa = false

	a.db.Save(&o)

	return o, true
}

// =========================================================
// SEEDS
// =========================================================

// SembrarSiVacio inserta datos iniciales solo si aún no hay datos
func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64

	a.db.Model(&modelos.Parqueadero{}).Count(&n)

	if n > 0 {
		return
	}

	// =========================
	// PARQUEADEROS
	// =========================

	parqueaderos := []modelos.Parqueadero{
		{ID: "1", Nombre: "Parqueadero FCVT", Ubicacion: "Facultad de Ciencias de la Vida y Tecnologías", Capacidad: 20, Activo: true},
		{ID: "2", Nombre: "Parqueadero Central", Ubicacion: "Centro de la ciudad", Capacidad: 50, Activo: true},
		{ID: "3", Nombre: "Parqueadero Norte", Ubicacion: "Zona norte de la ciudad", Capacidad: 30, Activo: true},
		{ID: "4", Nombre: "Parqueadero Sur", Ubicacion: "Zona sur de la ciudad", Capacidad: 25, Activo: true},
		{ID: "5", Nombre: "Parqueadero Motos", Ubicacion: "Entrada Principal", Capacidad: 60, Activo: false},
	}

	a.db.Create(&parqueaderos)

	// =========================
	// ESPACIOS
	// =========================

	espacios := []modelos.Espacio{
		{ID: "1", ParqueaderoID: "1", Numero: "A1", Disponible: false},
		{ID: "2", ParqueaderoID: "1", Numero: "A2", Disponible: false},
		{ID: "3", ParqueaderoID: "1", Numero: "A3", Disponible: true},
		{ID: "4", ParqueaderoID: "2", Numero: "B1", Disponible: false},
		{ID: "5", ParqueaderoID: "2", Numero: "B2", Disponible: true},
		{ID: "6", ParqueaderoID: "3", Numero: "C1", Disponible: true},
	}

	a.db.Create(&espacios)

	// =========================
	// OCUPACIONES
	// =========================

	ocupaciones := []modelos.Ocupacion{
		{
			ID:        1,
			EspacioID: "1",
			Placa:     "ABC123",
			Entrada:   time.Now().Add(-2 * time.Hour),
			Activa:    true,
		},
		{
			ID:        2,
			EspacioID: "2",
			Placa:     "XYZ456",
			Entrada:   time.Now().Add(-1 * time.Hour),
			Activa:    true,
		},
		{
			ID:        3,
			EspacioID: "4",
			Placa:     "ULE789",
			Entrada:   time.Now().Add(-30 * time.Minute),
			Activa:    true,
		},
	}

	a.db.Create(&ocupaciones)
}

// Chequeo en tiempo de compilación: AlmacenSQLite debe cumplir Almacen.
var _ Almacen = (*AlmacenSQLite)(nil)
