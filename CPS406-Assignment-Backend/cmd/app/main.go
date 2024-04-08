package main

import (
	"CPS406-Assignment-Backend/internal/api/http/server"
	"CPS406-Assignment-Backend/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// main is the entry point of the application
func main() {
	// Connect to the database
	dataBase := db.ConnectDB()
	// Migrate the database
	db.MigrateDB(dataBase)
	// Seed the database
	db.SeedDatabase(dataBase)
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// A good base middleware stack
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	// Routes for the API
	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			_, err := w.Write([]byte("The server is running!"))
			if err != nil {
				return
			}
		})
		server.Server(r, dataBase)
	})
	// Serve the Vue app
	dist := os.Getenv("DIST")
	serveVueApp(r, dist)
	// Listen for requests on port in your .env file
	portNum := ":" + os.Getenv("PORT")
	err = http.ListenAndServe(portNum, r)
	if err != nil {
		return
	}
}

// serveVueApp serves the Vue app
func serveVueApp(r *chi.Mux, fsRoot string) {
	// Find the absolute path of the Vue app
	absPath, err := filepath.Abs(fsRoot)
	if err != nil {
		log.Fatalf("Error calculating absolute path: %s", err)
	}
	// File server for the Vue app
	fs := http.FileServer(http.Dir(absPath))
	// Handle all requests to the root URL from the dist folder
	r.Handle("/*", http.StripPrefix("/", fs)) // Serve static files
}
