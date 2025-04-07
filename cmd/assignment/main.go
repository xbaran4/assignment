package main

import (
	"assignment/pkg/repository"
	"assignment/pkg/rest"
	"errors"
	"log"
	"net/http"
	"os"
)

func main() {
	dbFilename := getEnv("DB_FILENAME", "gorm.db")
	appPort := getEnv("APP_PORT", "8080")

	userHandler := rest.UserHandler{Repo: repository.InitUserRepository(dbFilename)}
	server := rest.SetupServer(userHandler, appPort)

	log.Println("starting server on port " + appPort)
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("error starting server: %s\n", err)
	}
	log.Println("server closed")
}

func getEnv(key, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}
