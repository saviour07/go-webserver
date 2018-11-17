package controller

import (
	"github.com/gorilla/mux"
	"github.com/saviour07/go-webserver/handlers"
	"github.com/saviour07/go-webserver/models"
)

// HomeController structure
type HomeController struct {
	handler handlers.HomeHandler
}

// Path implements controller.Path()
func (HomeController) Path() string {
	return "/"
}

// Title implements controller.Title()
func (HomeController) Title() string {
	return "Home"
}

// NavItem implements controller.NavItem()
func (h HomeController) NavItem() models.NavItem {
	return models.NavItem{
		Path:     h.Path(),
		Title:    h.Title(),
		Selected: true,
	}
}

// RegisterHandler implements handlers.RegisterHandler()
func (h HomeController) RegisterHandler(navBar []models.NavItem, router *mux.Router) {
	h.handler.RegisterHandler(h.Path(), navBar, router)
}
