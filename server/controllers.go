package main

import "net/http"

func (s *APIServer) handleRoute(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusAccepted, "Hello From PUPPY API")
}
