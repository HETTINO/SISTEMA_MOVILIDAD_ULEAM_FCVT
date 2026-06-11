package main

import (
	"fmt"
	"log"
	"net/http"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/handlers"
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/middleware"
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/modelos"
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/storage"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func main() {
	// 1. Abrir SQLite y migrar el esquema del equipo
	db, err := gorm.Open(sqlite.Open("parking.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}

	if err := db.AutoMigrate(
		&modelos.Parqueadero{},
		&modelos.Espacio{},
		&modelos.Ocupacion{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// 2. Crear el almacenamiento SQLite y sembrar si está vacío
	parkingStore := storage.NuevoAlmacenSQLite(db)
	parkingStore.SembrarSiVacio()

	// === NUEVO: Inicializar tu almacenamiento en Memoria compartido ===
	memStore := storage.NewMemoria()
	memStore.SeedModuloAcceso() // <--- ¡AUTOMATIZACIÓN AGREGADA AQUÍ!

	// 3. Inicializar handlers (Los del grupo + el tuyo centralizado)
	parkingHandler := handlers.NewParkingHandler(parkingStore)
	accesoHandler := handlers.NewAccesoHandler(memStore)

	// 4. Router
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// ==========================================
	// RUTAS DE PARQUEADEROS (EQUIPO)
	// ==========================================
	r.Route("/api/v1/parking", func(r chi.Router) {
		r.Get("/", parkingHandler.ListarParqueaderos)
		r.Get("/{id}", parkingHandler.ObtenerParqueadero)
		r.Post("/", parkingHandler.CrearParqueadero)
		r.Put("/{id}", parkingHandler.ActualizarParqueadero)
		r.Delete("/{id}", parkingHandler.EliminarParqueadero)
	})

	// ==========================================
	// RUTAS DE ESPACIOS (EQUIPO)
	// ==========================================
	r.Route("/api/v1/espacios", func(r chi.Router) {
		r.Get("/", parkingHandler.ListarEspacios)
		r.Get("/{id}", parkingHandler.ObtenerEspacio)
		r.Post("/", parkingHandler.CrearEspacio)
		r.Put("/{id}", parkingHandler.ActualizarEspacio)
		r.Delete("/{id}", parkingHandler.EliminarEspacio)
	})

	// ==========================================
	// RUTAS DE OCUPACIONES (EQUIPO)
	// ==========================================
	r.Route("/api/v1/ocupaciones", func(r chi.Router) {
		r.Get("/", parkingHandler.ListarOcupaciones)
		r.Get("/{id}", parkingHandler.ObtenerOcupacion)
		r.Post("/", parkingHandler.CrearOcupacion)
		r.Put("/{id}/liberar", parkingHandler.LiberarOcupacion)
	})

	// =======================================================
	// === RUTAS DE TU MÓDULO: ACCESO DE ENTRADA Y SALIDA ===
	// =======================================================

	// 1. Entidad: Usuarios
	r.Route("/api/v1/usuarios", func(r chi.Router) {
		r.Get("/", accesoHandler.ListarUsuarios)
		r.Get("/{id}", accesoHandler.ObtenerUsuario)
		r.Post("/", accesoHandler.CrearUsuario)
		r.Put("/{id}", accesoHandler.ActualizarUsuario)
		r.Delete("/{id}", accesoHandler.EliminarUsuario)
	})

	// 2. Entidad: Vehículos
	r.Route("/api/v1/vehiculos", func(r chi.Router) {
		r.Get("/", accesoHandler.ListarVehiculos)
		r.Get("/{id}", accesoHandler.ObtenerVehiculo)
		r.Post("/", accesoHandler.CrearVehiculo)
		r.Put("/{id}", accesoHandler.ActualizarVehiculo)
		r.Delete("/{id}", accesoHandler.EliminarVehiculo)
	})

	// 3. Entidad: Puntos de Acceso (Garitas)
	r.Route("/api/v1/puntos-acceso", func(r chi.Router) {
		r.Get("/", accesoHandler.ListarPuntosAcceso)
		r.Get("/{id}", accesoHandler.ObtenerPuntoAcceso)
		r.Post("/", accesoHandler.CrearPuntoAcceso)
		r.Put("/{id}", accesoHandler.ActualizarPuntoAcceso)
		r.Delete("/{id}", accesoHandler.EliminarPuntoAcceso)
	})

	// 4. Entidad: Control de Accesos (Transaccional)
	r.Route("/api/v1/acceso", func(r chi.Router) {
		r.Get("/", accesoHandler.ListarAccesos)
		r.Get("/{id}", accesoHandler.ObtenerAcceso)
		r.Post("/", accesoHandler.CrearAcceso)
		r.Put("/{id}", accesoHandler.ActualizarAcceso)
		r.Delete("/{id}", accesoHandler.EliminarAcceso)
	})

	// Mensajes de inicialización del sistema
	fmt.Println("===================================")
	fmt.Println(" SISTEMA MOVILIDAD ULEAM FCVT ")
	fmt.Println("===================================")
	fmt.Println("Servidor iniciado con éxito en:")
	fmt.Println("http://localhost:8080")
	fmt.Println("===================================")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
