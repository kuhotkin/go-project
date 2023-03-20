package handler

import (
	"github.com/kuhotkin/jira-connector/pkg/service"
	"net/http"
)

type Handler struct {
	service service.JiraConnector
}

func NewHandler(service service.JiraConnector) *Handler {
	return &Handler{
		service,
	}
}

func (hand *Handler) InitHandleFuncs() {
	http.HandleFunc("/projects", hand.GetProjects)

}
