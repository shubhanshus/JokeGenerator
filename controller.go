package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetJokeHandler returns the joke to the requester
func GetJokeHandler(w http.ResponseWriter, r *http.Request) {
	person, respCode := GetPerson()

	//Return error message when we got an invalid response from name server
	if respCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error! Please try again in sometime")
		log.Printf("Invalid response from Name API Server")
		return
	}
	joke, respCode := GetJokeFromAPI(person)

	//Return error message when we got an invalid response from Joke server
	if respCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error! Please try again in sometime")
		log.Printf("Invalid response from Joke API Server")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(joke)
	return
}

// HealthChecker Handler checks the status of server
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Server is running")
}
