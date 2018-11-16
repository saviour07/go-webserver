package controller

import (
	"github.com/gorilla/mux"
	"github.com/saviour07/go-webserver/models"
)

// Controller interface
type Controller interface {
	Path() string
	Title() string
	NavItem() models.NavItem
	RegisterHandler(navBar []models.NavItem, router *mux.Router)
}
