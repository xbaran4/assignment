package rest

import (
	"assignment/pkg/dto"
	"assignment/pkg/repository"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Repo repository.UserRepository
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.IDlessUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("could not parse user from request body: %s", err), http.StatusBadRequest)
		return
	}

	persistedUser, err := h.Repo.CreateUser(user.ToModel())
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to store user: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(dto.FromModelWithID(persistedUser)); err != nil {
		log.Printf("failed to write http response: %s", err)
	}
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idPathValue := r.PathValue("id")
	id, err := strconv.ParseUint(idPathValue, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user, err := h.Repo.GetUser(uint(id))
	if errors.Is(err, repository.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(dto.FromModelWithoutID(user)); err != nil {
		log.Printf("failed to write http response: %s", err)
	}

}
