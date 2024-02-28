package main

import "net/http"

type APIServer struct {
	listenAddress string
	// store      Storage
}

type apiFunc func(http.ResponseWriter, *http.Request) error
type ApiError struct {
	Error string `json:"error"`
}

// GET STORE FROM DB
// store, err := NewPostgresStore()
// if err != nil {
// log.Fatal(err)
// }
// if err := store.Init(); err != nil {
// log.Fatal(err)
// }
// RUNNING API SERVER FROM API.GO
