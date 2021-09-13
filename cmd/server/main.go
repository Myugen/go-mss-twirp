package main

import (
	"net/http"

	"github.com/myugen/go-mss-twirp/internal/user_server"
	"github.com/myugen/go-mss-twirp/rpc/user"
)

func main() {
	server := user_server.New() // implements Haberdasher interface
	twirpHandler := user.NewUserServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
