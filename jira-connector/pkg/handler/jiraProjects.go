package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (handler *Handler) GetProjects(w http.ResponseWriter, req *http.Request) {
	details, err := handler.service.Get(req.Context())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Printf("Encode error: %v\n", err)
	}
}
