package main

// Joke represents a Joke JSON response from icndb API
type Joke struct {
	Type  string `json:"type"`
	Value struct {
		ID   int    `json:"id"`
		Joke string `json:"joke"`
	} `json:"value"`
}

// Person represents a Person JSON response from uinames API
type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}
