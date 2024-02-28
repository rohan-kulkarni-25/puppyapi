// UTILS FUNCTIONS
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

// func WriteImage(w http.ResponseWriter, status int, v []byte) error {
// 	w.Header().Add("Content-Type", "application/octet-stream")
// 	w.WriteHeader(status)
// 	w.Write(v)
// 	return "";
// }

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func getParams(r *http.Request, paramName string) (string, error) {
	param := mux.Vars(r)[paramName]
	return param, nil
}
