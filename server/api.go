package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NEWAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/", makeHTTPHandleFunc(s.handleRoute))
	router.HandleFunc("/dogs/{name}", makeHTTPHandleFunc(s.handleGetDogs))
	router.HandleFunc("/dogs", makeHTTPHandleFunc(s.handleGetDogs))
	log.Println("PUPPY API server running on port: ", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}
