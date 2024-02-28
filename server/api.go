package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NEWAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/", makeHTTPHandleFunc(s.handleRoute))
	log.Println("PUPPY API server running on port: ", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}
