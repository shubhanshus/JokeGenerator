package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Get the Person JSON
func GetPerson() (Person, int) {
	// Fill the record with the data from the JSON
	var record Person

	nameApiURL := "http://uinames.com/api/"

	// Build the request
	req, err := http.NewRequest("GET", nameApiURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return record, http.StatusInternalServerError
	}

	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return record, http.StatusInternalServerError
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// check for status code
	if resp.StatusCode == http.StatusOK {
		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
			log.Println("JSON Name fetch", err)
		}
	} else {
		log.Printf("Received Status: %d instead of 200", resp.StatusCode)
	}

	return record, resp.StatusCode
}

// Get the Joke JSON
func GetJokeFromAPI(person Person) (string, int) {
	// QueryEscape escapes the FirstName and LastName string so
	// it can be safely placed inside a URL query
	safeFirstName := url.QueryEscape(person.Name)
	safeSurname := url.QueryEscape(person.Surname)

	url := fmt.Sprintf("http://api.icndb.com/jokes/random?firstName=%s&lastName=%s&limitTo=\\[nerdy\\]", safeFirstName, safeSurname)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return "", http.StatusInternalServerError
	}

	// create a HTTP client
	client := &http.Client{}

	// Send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return "", http.StatusInternalServerError
	}

	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Joke

	// check for status code
	if resp.StatusCode == http.StatusOK {
		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
			log.Println("JSON Name fetch error", err)
			return "", http.StatusInternalServerError
		}
	} else {
		log.Printf("Received Status: %d instead of 200", resp.StatusCode)
	}
	return record.Value.Joke, http.StatusOK

}
