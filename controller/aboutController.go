package controller

import (
	"github.com/gorilla/mux"
	"github.com/saviour07/go-webserver/handlers"
	"github.com/saviour07/go-webserver/models"
)

// AboutController structure
type AboutController struct {
	handlers.AboutHandler
}

// Path implements controller.Path()
func (AboutController) Path() string {
	return "/about"
}

// Title implements controller.Title()
func (AboutController) Title() string {
	return "About"
}

// NavItem implements controller.NavItem()
func (a *AboutController) NavItem() models.NavItem {
	return models.NavItem{
		Path:     a.Path(),
		Title:    a.Title(),
		Selected: false,
	}
}

// RegisterHandler implements controller.RegisterHandler()
func (a *AboutController) RegisterHandler(navBar models.NavigationItems, router *mux.Router) {
	a.Register(a.Path(), navBar, router)
}
