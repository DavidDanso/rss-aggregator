package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")

	router := chi.NewRouter()

	server := &http.Server{
		Addr: ":" + portString,
		Handler: router,
	}
	
	log.Printf("Starting server on port %s", portString)
	err := server.ListenAndServe()
	if err != nil {
        log.Fatal(err)
    }

}