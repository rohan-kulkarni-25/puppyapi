package main

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type APIServer struct {
	listenAddress string
	store         Storage
}

type apiFunc func(http.ResponseWriter, *http.Request) error
type ApiError struct {
	Error string `json:"error"`
}

type DogResponse struct {
	Param string `json:"param"`
	Data  []byte `json:"data"`
}

type Storage interface {
	createDogsTable() error
	GetDogImageByName(string) (*Dog, error)
	AddDog(dog *Dog) error
}

type Dog struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type PostgresStore struct {
	db sqlx.DB
}

func NewDog(name string, path string) *Dog {
	return &Dog{
		Name: name, Path: path,
	}
}
