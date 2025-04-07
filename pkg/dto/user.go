package dto

import (
	"assignment/pkg/repository"
	"time"
)

type IDlessUser struct {
	ExternalId  string    `json:"external_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

func (u IDlessUser) ToModel() repository.User {
	return repository.User{
		ExternalId:  u.ExternalId,
		Name:        u.Name,
		Email:       u.Email,
		DateOfBirth: u.DateOfBirth,
	}
}

func FromModelWithoutID(user repository.User) IDlessUser {
	return IDlessUser{
		ExternalId:  user.ExternalId,
		Name:        user.Name,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
	}
}

type IDiedUser struct {
	ID          uint      `json:"id"`
	ExternalId  string    `json:"external_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

func (u IDiedUser) ToModel() repository.User {
	return repository.User{
		ID:          u.ID,
		ExternalId:  u.ExternalId,
		Name:        u.Name,
		Email:       u.Email,
		DateOfBirth: u.DateOfBirth,
	}
}

func FromModelWithID(user repository.User) IDiedUser {
	return IDiedUser{
		ID:          user.ID,
		ExternalId:  user.ExternalId,
		Name:        user.Name,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
	}
}
