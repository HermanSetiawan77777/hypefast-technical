package httpserver

import (
	"hypefast-technical/internal/httpserver/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/shorten", handler.ShortenUrl).Methods(http.MethodPost)
	r.HandleFunc("/{id}/stats", handler.GetLinkStats).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handler.GetURL).Methods(http.MethodGet)

	return r
}
