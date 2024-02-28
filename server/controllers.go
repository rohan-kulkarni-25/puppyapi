package main

import (
	"encoding/base64"
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
	var bytesData []byte
	for _, val := range fileBytes {
		bytesData = append(bytesData, byte(val))
	}

	// Encode to base64
	base64Encoded := base64.StdEncoding.EncodeToString(bytesData)
	// response := DogResponse{
	// 	Param: dogs.Name,
	// 	Data:  []byte(base64Encoded),
	// }
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(base64Encoded))
	return err
	// return WriteJSON(w, http.StatusAccepted, response)
}

// DOG BREEDS
// labrador
// pitbull
// dachshund
// african
// affenpinscher
