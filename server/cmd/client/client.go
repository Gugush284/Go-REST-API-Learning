package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	MakeRequest()
}

func MakeRequest() {

	message := map[string]interface{}{
		"login":    "Slava8924",
		"password": "JGjkgkggd9",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
}
