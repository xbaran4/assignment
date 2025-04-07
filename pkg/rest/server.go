package rest

import (
	"net/http"
)

func SetupServer(userHandler UserHandler, port string) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("POST /save", http.HandlerFunc(userHandler.CreateUser))
	mux.Handle("GET /{id}", http.HandlerFunc(userHandler.GetUser))
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	return srv
}
