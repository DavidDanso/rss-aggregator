package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")

	r := chi.NewRouter()

	// CORS configuration (replace with your allowed origins, methods, etc.)
    cors := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://*", "https://*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Authorization", "Content-Type"},
        ExposedHeaders:   []string{"Content-Length", "Accept"},
        AllowCredentials: false,
		MaxAge:           300,
    })

    // Apply CORS middleware to all routes
    r.Use(cors.Handler)

	v1Router := chi.NewRouter()
	v1Router. Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	r.Mount("/v1", v1Router)


	server := &http.Server{
		Addr: ":" + portString,
		Handler: r,
	}
	
	log.Printf("Starting server on port %s", portString)
	err := server.ListenAndServe()
	if err != nil {
        log.Fatal(err)
    }

}