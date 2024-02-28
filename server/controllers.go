package main

import (
	"fmt"
	"net/http"
	"os"
)

func (s *APIServer) handleRoute(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusAccepted, "Hello From PUPPY API")
}

func (s *APIServer) handleGetDogs(w http.ResponseWriter, r *http.Request) error {
	param, _ := getParams(r, "name")
	dogs, err := s.store.GetDogImageByName(param)

	if err != nil {
		return err
	}
	fmt.Println(dogs)
	fileBytes, err := os.ReadFile(dogs.Path)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}
	response := DogResponse{
		Param: dogs.Name,
		Data:  fileBytes,
	}
	return WriteJSON(w, http.StatusAccepted, response)
}

// DOG BREEDS
// labrador
// pitbull
// dachshund
// african
// affenpinscher
