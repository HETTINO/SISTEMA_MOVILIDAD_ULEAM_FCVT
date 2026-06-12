package main

import (
	"fmt"
	"log"
	"net/http"

	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/handlers"
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/middleware"
	"SISTEMA_MOVILIDAD_ULEAM_FCVT/internal/storage"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func main() {
	// 1. Inicializar Almacenamiento (Memoria)
	memStore := storage.NewMemoria()

	// 2. Cargar Seeds del Módulo de Transporte
	memStore.SeedRutas()
	memStore.SeedParadas()
	memStore.SeedCarritos()
	memStore.SeedLocaciones()
	memStore.SeedSolicitudes()

	// 3. Inicializar Handlers
	transporteHandler := handlers.NewTransporteHandler(memStore)

	// 4. Router
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// Montar rutas del módulo de transporte
	r.Mount("/api/v1/transporte", transporteHandler.Rutas())

	fmt.Println("=======================================")
	fmt.Println(" SISTEMA MOVILIDAD ULEAM - TRANSPORTE ")
	fmt.Println(" Puerto: 8080                          ")
	fmt.Println("=======================================")

	log.Fatal(http.ListenAndServe(":8080", r))
}