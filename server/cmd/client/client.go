package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	Menu()

	for {
		var key int

		fmt.Fprint(os.Stdout, "Choose func\n")
		fmt.Fscan(os.Stdin, &key)

		switch key {
		case 0:
			os.Exit(0)
		case 1:
			CreateRequest()
		case 2:
			SessionRequest()
		case 3:
			Menu()
		default:
			fmt.Fprint(os.Stdout, "Unknown case\n")
		}
	}
}

func Menu() {
	fmt.Fprint(os.Stdout, "Menu:\n")
	fmt.Fprint(os.Stdout, " case 0: Exit\n")
	fmt.Fprint(os.Stdout, " case 1: CreateRequest\n")
	fmt.Fprint(os.Stdout, " case 2: SessionRequest\n")
	fmt.Fprint(os.Stdout, " case 3: Menu\n")
}

func SessionRequest() {

	message := map[string]interface{}{
		"login":    "Slava8924",
		"password": "JGjkgkggd9",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:8080/sessions", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
	log.Println(resp.Cookies())
}

func CreateRequest() {

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
	log.Println(resp.Cookies())
}
