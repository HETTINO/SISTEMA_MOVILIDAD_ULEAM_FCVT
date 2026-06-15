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
	db, err := gorm.Open(sqlite.Open("parking.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&modelos.Parqueadero{}, &modelos.Espacio{}, &modelos.Ocupacion{})

	// Inicialización: Si 'NuevoAlmacenSQLite' te da error,
	// cámbialo por 'storage.NewMemoria()' o el constructor real que tengan tus compañeros.
	store := storage.NewMemoria()
	store.SeedTransporteInterno()

	// Handlers
	parkingHandler := handlers.NewParkingHandler(store)
	transporteHandler := handlers.NewTransporteHandler(store)

	r := chi.NewRouter()
	r.Use(chimw.CleanPath, chimw.Logger, chimw.Recoverer, middleware.CORS)

	// Rutas Parqueaderos
	r.Route("/api/v1/parking", func(r chi.Router) {
		r.Get("/", parkingHandler.ListarParqueaderos)
		r.Get("/{id}", parkingHandler.ObtenerParqueadero)
		r.Post("/", parkingHandler.CrearParqueadero)
		r.Put("/{id}", parkingHandler.ActualizarParqueadero)
		r.Delete("/{id}", parkingHandler.EliminarParqueadero)
	})

	// Rutas Transporte
	r.Route("/api/v1/transporte", func(r chi.Router) {
		r.Get("/solicitudes", transporteHandler.ListarSolicitudes)
		r.Get("/solicitudes/{id}", transporteHandler.ObtenerSolicitud)
		r.Post("/solicitudes", transporteHandler.CrearSolicitud)
		r.Put("/solicitudes/{id}", transporteHandler.ActualizarSolicitud)
		r.Delete("/solicitudes/{id}", transporteHandler.EliminarSolicitud)
	}) // <-- CORREGIDO: Aquí cerrabas con un ) y debía ser }

	fmt.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
