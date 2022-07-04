package main

import (
	"hypefast-technical/internal/httpserver"
	"net/http"
)

func main() {
	//host should be in env so it will automatically update when the host is changed
	appserver := &http.Server{
		Addr:    "localhost:8080",
		Handler: httpserver.HandleRoutes(),
	}

	appserver.ListenAndServe()
}
