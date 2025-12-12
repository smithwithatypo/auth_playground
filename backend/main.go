package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var dbURL string
	env := os.Getenv("environment")

	if env == "production" {
		dbURL = os.Getenv("DATABASE_URL") // Railway's private network
	} else {
		dbURL = os.Getenv("DATABASE_PUBLIC_URL") // local dev
	}

	if dbURL == "" {
		log.Fatal("no database URL configured")
	}

	var err error

	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	log.Println("Connected to database")

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
