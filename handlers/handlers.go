package handlers

import (
	"github.com/gorilla/mux"
	"github.com/saviour07/go-webserver/models"
)

// Handler interface
type Handler interface {
	RegisterHandler(path string, navItems []models.NavItem, router *mux.Router)
}
