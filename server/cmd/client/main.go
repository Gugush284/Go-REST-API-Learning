package main

import (
	"log"
	"net/http"

	ServerClient "github.com/Gugush284/Go-server.git/internal/client"
)

func main() {
	var cookie []*http.Cookie

	log.Println("CreateRequest")
	ServerClient.CreateRequest()

	log.Println("SessionRequest")
	cookie = ServerClient.SessionRequest()

	log.Println("WhoamiRequest")
	ServerClient.WhoamiRequest(cookie)
}
