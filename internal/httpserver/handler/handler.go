package handler

import (
	"encoding/json"
	"fmt"
	"hypefast-technical/internal/service/link"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func ShortenUrl(w http.ResponseWriter, r *http.Request) {
	type shortenPayload struct {
		Url         string `json:"url"`
		OptionShort string `json:"OptionShort"`
	}

	var p *shortenPayload
	err := json.NewDecoder(r.Body).Decode(&p)
	response := map[string]interface{}{}
	if err != nil {
		if err == io.EOF {
			response["error"] = "Please fill payload"
			responseWithData(w, http.StatusBadRequest, response)
			return
		}
		response["error"] = "Error parsing payload"
		responseWithData(w, http.StatusInternalServerError, response)
		return
	}

	if p.Url == "" {
		response["error"] = "Please fill url"
		responseWithData(w, http.StatusInternalServerError, response)
		return
	}

	newLink := link.AddNewLink(p.Url, p.OptionShort)

	//host should be in env so it will automatically update when the host is changed
	response["link"] = fmt.Sprintf("http://localhost:8080/%s", newLink.Id)
	responseWithData(w, http.StatusOK, response)
}

func GetURL(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	existingLink := link.GetLinkByID(id)
	response := map[string]interface{}{}
	if existingLink == nil || existingLink.Url == "" {
		response["error"] = "Link not found"
		responseWithData(w, http.StatusNotFound, response)
		return
	}
	err := link.UpdateRedirectCount(existingLink.Id, existingLink.RedirectCount+1)
	if err != nil {
		response["error"] = "Link not found"
		responseWithData(w, http.StatusNotFound, response)
		return
	}

	http.Redirect(w, r, existingLink.Url, http.StatusPermanentRedirect)
}

func GetLinkStats(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	existingLink := link.GetLinkByID(id)
	response := map[string]interface{}{}
	if existingLink == nil || existingLink.Url == "" {
		response["error"] = "Link not found"
		responseWithData(w, http.StatusNotFound, response)
		return
	}

	response["redirect_count"] = existingLink.RedirectCount
	response["created_at"] = existingLink.CreatedAt
	responseWithData(w, http.StatusOK, response)
}

func responseWithData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
