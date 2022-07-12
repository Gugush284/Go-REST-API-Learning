package main

import (
	"net/http"

	ServerClient "github.com/Gugush284/Go-server.git/internal/client"
)

func main() {
	var cookie []*http.Cookie

	ServerClient.CreateRequest()
	cookie = ServerClient.SessionRequest()
	ServerClient.WhoamiRequest(cookie)
}
