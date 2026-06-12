package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
	// 1. Abrir SQLite y migrar el esquema
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

	// 3. Inicializar handlers
	parkingHandler := handlers.NewParkingHandler(parkingStore)

	// 4. Router
	r := chi.NewRouter()

	// ALimpia barras dobles y elimina la barra final (trailing slash) automáticamente
	r.Use(chimw.CleanPath)

	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// =========================
	// RUTAS DE PARQUEADEROS
	// =========================
	r.Route("/api/v1/parking", func(r chi.Router) {
		r.Get("/", parkingHandler.ListarParqueaderos)
		r.Get("/{id}", parkingHandler.ObtenerParqueadero)
		r.Post("/", parkingHandler.CrearParqueadero)
		r.Put("/{id}", parkingHandler.ActualizarParqueadero)
		r.Delete("/{id}", parkingHandler.EliminarParqueadero)
	})

	// =========================
	// RUTAS DE ESPACIOS
	// =========================
	r.Route("/api/v1/espacios", func(r chi.Router) {
		r.Get("/", parkingHandler.ListarEspacios)
		r.Get("/{id}", parkingHandler.ObtenerEspacio)

		r.Post("/", parkingHandler.CrearEspacio)
		r.Put("/{id}", parkingHandler.ActualizarEspacio)
		r.Delete("/{id}", parkingHandler.EliminarEspacio)
	})

	// =========================
	// RUTAS DE OCUPACIONES
	// =========================
	r.Route("/api/v1/ocupaciones", func(r chi.Router) {
		r.Get("/", parkingHandler.ListarOcupaciones)
		r.Get("/{id}", parkingHandler.ObtenerOcupacion)

		r.Post("/", parkingHandler.CrearOcupacion)

		r.Put("/{id}/liberar", parkingHandler.LiberarOcupacion)
	})

	// =========================
	// RUTAS DE ACCESO / LOGIN (Mensaje Automático)
	// =========================
	r.Route("/api/v1/acceso", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			// Genera un número aleatorio de 4 dígitos para la placa
			numeroAleatorio := time.Now().UnixNano()%9000 + 1000
			placaAutomatica := fmt.Sprintf("ABC-%d", numeroAleatorio)

			// Crea la respuesta automática en una sola línea
			respuesta := map[string]interface{}{
				"placa_vehiculo":  placaAutomatica,
				"id_punto_acceso": 1,
				"estado":          "Activo",
				"observaciones":   "Ingreso automático generado por el servidor",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(respuesta)
		})
	})

	fmt.Println("===================================")
	fmt.Println(" SISTEMA MOVILIDAD ULEAM FCVT ")
	fmt.Println("===================================")
	fmt.Println("Servidor iniciado en:")
	fmt.Println("http://localhost:8080")
	fmt.Println("===================================")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
