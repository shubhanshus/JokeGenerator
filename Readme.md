# Random Joke Generator

This project contains a web service which combines two existing web services to generate a Joke.
1) Fetch a random name from http://uinames.com/api/ 
2) Fetch a random Chuck Norris joke from http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy] 
3) Combine the results and return them to the user.


### Example 
Fetching a name
``` 
$ curl http://uinames.com/api/                                                                        
{"name":"Δαμέας","surname":"Γιάνναρης","gender":"male","region":"Greece"}
``` 
Fetching a joke
``` 
$ curl 'http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=\[nerdy\]'
{ "type": "success", "value": { "id": 181, "joke": "John Doe's OSI network model has only one layer - Physical.", "categories": [“nerdy”] } }
``` 
Using the new web service
``` 
$ curl ‘http://localhost:5000’
Δαμέας Γιάνναρης’s OSI network model has only one layer - Physical..
``` 

### Tools Needed:
1) Golang 
2) Dep 

### Steps to run the project:
1) To get all the dependencies type "dep ensure" in the command line from project folder 
2) Type "go build" in the command line from the project folder
3) Above command will create an executable which you can run with command ./JokeGenerator
4) The above command will start the server on port 8080
5) If you want to change the port just type export PORT=8000 on the command line before running the executable


### Project File Description
1) controller.go: This file contains the Handler functions for this web service
2) helpers.go: This file contains helper functions to get the name and joke from the name and joke api
3) main.go: main file which contains the main function for running the web server
4) main_test.go: Test file which tests the different handlers
5) model.go: This file contains models for Person and Joke
