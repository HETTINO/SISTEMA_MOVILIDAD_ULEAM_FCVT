package main

import (
	"fmt"
	"log"
	"net/http"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/handlers"
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Inicializar stores
	parkingStore := storage.NewParkingStore()

	// Inicializar handlers
	parkingHandler := handlers.NewParkingHandler(parkingStore)

	// Configurar router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Rutas del módulo de parqueo (Eduardo - Grupo F)
	r.Route("/api/v1/parking", func(r chi.Router) {
		// CRUD de Parqueaderos
		r.Get("/", parkingHandler.ListarParqueaderos)         // Read-todos
		r.Post("/", parkingHandler.CrearParqueadero)          // Create
		r.Get("/{id}", parkingHandler.ObtenerParqueadero)     // Read-uno
		r.Put("/{id}", parkingHandler.ActualizarParqueadero)  // Update
		r.Delete("/{id}", parkingHandler.EliminarParqueadero) // Delete

		// Sub-rutas por parqueadero
		r.Get("/{id}/espacios", parkingHandler.ListarEspacios)
		r.Get("/{id}/disponibilidad", parkingHandler.ObtenerDisponibilidad)
		r.Get("/{id}/status", parkingHandler.ObtenerEstado)
	})

	r.Route("/api/v1/parking/espacios", func(r chi.Router) {
		// CRUD de Espacios
		r.Post("/", parkingHandler.CrearEspacio)                         // Create
		r.Get("/{espacioID}", parkingHandler.ObtenerEspacio)             // Read-uno
		r.Patch("/{espacioID}/reserve", parkingHandler.ReservarEspacio)  // reservar
		r.Patch("/{espacioID}/release", parkingHandler.LiberarEspacio)   // liberar
		r.Delete("/{espacioID}", parkingHandler.EliminarEspacio)         // Delete
		r.Post("/{espacioID}/ocupar", parkingHandler.RegistrarOcupacion) // registrar entrada
	})

	r.Route("/api/v1/parking/ocupacion", func(r chi.Router) {
		r.Get("/", parkingHandler.ListarOcupaciones)
	})

	fmt.Println("Servidor corriendo en http://localhost:8080")
	fmt.Println("Módulo de Parqueo - Eduardo López")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error al iniciar servidor:", err)
	}
}
