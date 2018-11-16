package handlers

import (
	"github.com/gorilla/mux"
	"github.com/saviour07/go-webserver/models"
)

// Handler interface
type Handler interface {
	AddNavBar(navItems []models.NavItem)
	RegisterHandler(path string, router *mux.Router)
}
