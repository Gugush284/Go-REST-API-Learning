package main

import (
	"log"
	"net/http"

	ServerClient "github.com/Gugush284/Go-server.git/third_party/client/internal"
)

func main() {
	var cookie []*http.Cookie

	log.Println("CreateRequest")
	ServerClient.CreateRequest()

	log.Println("SessionRequest")
	cookie = ServerClient.SessionRequest()

	log.Println("WhoamiRequest")
	ServerClient.WhoamiRequest(cookie)

	log.Println("Upload")
	id := ServerClient.Upload(cookie)
	if id != 0 {
		log.Println("Download")
		ServerClient.Download(id)
	}
}
