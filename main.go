package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {

	// Get the port variable from environment variables
	port := os.Getenv("PORT")

	// if port is not set use the default port of 8080
	if port == "" {
		port = "8080"
	}
	// Create New MuxRouter
	router := mux.NewRouter()

	// Enable CORS requests
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	handler := c.Handler(router)
	router.HandleFunc("/", GetJokeHandler).Methods("GET")
	router.HandleFunc("/healthcheck", healthCheckHandler).Methods("GET")

	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Server is now running on port ", port)
	}
}
