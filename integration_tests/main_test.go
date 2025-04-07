package integration_tests

import (
	"assignment/pkg/repository"
	"assignment/pkg/rest"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestGetNonexistentUser(t *testing.T) {
	userHandler := rest.UserHandler{Repo: repository.InitUserRepository(":memory:")}
	s := rest.SetupServer(userHandler, "8000")
	go s.ListenAndServe()
	defer s.Close()

	resp, err := http.Get("http://localhost:8000/0")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
	}
}

func TestGetUser(t *testing.T) {
	repo := repository.InitUserRepository(":memory:")
	_, err := repo.CreateUser(repository.User{
		ID:          5,
		ExternalId:  "abc",
		Name:        "efg",
		Email:       "ejk",
		DateOfBirth: time.Time{},
	})
	if err != nil {
		t.Fatal(err)
	}
	userHandler := rest.UserHandler{Repo: repo}
	s := rest.SetupServer(userHandler, "8000")
	go s.ListenAndServe()
	defer s.Close()

	resp, err := http.Get("http://localhost:8000/5")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestCreateUser(t *testing.T) {
	userHandler := rest.UserHandler{Repo: repository.InitUserRepository(":memory:")}
	s := rest.SetupServer(userHandler, "8000")
	go s.ListenAndServe()
	defer s.Close()

	user := `{
	"external_id": "31178118-b1f7-4b6c-ab21-818b92b56714",
	"name": "some name",
	"email": "email@email.com",
	"date_of_birth": "2020-01-01T12:12:34+00:00"
}`
	reader := strings.NewReader(user)
	resp, err := http.Post("http://localhost:8000/save", "application/json", reader)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestCreateUserInvalidRequestBody(t *testing.T) {
	userHandler := rest.UserHandler{Repo: repository.InitUserRepository(":memory:")}
	s := rest.SetupServer(userHandler, "8000")
	go s.ListenAndServe()
	defer s.Close()

	user := `{"name": 5}`
	reader := strings.NewReader(user)
	resp, err := http.Post("http://localhost:8000/save", "application/json", reader)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}
