package controller

import (
	"github.com/gorilla/mux"
	"github.com/saviour07/go-webserver/handlers"
	"github.com/saviour07/go-webserver/models"
)

// AccordionController structure
type AccordionController struct {
	handlers.AccordionHandler
}

// Path implements controller.Path()
func (AccordionController) Path() string {
	return "/accordion"
}

// Title implements controller.Title()
func (AccordionController) Title() string {
	return "Accordion"
}

// NavItem implements controller.NavItem()
func (h *AccordionController) NavItem() models.NavItem {
	return models.NavItem{
		Path:     h.Path(),
		Title:    h.Title(),
		Selected: false,
	}
}

// RegisterHandler implements controller.RegisterHandler()
func (h *AccordionController) RegisterHandler(navBar models.NavigationItems, router *mux.Router) {
	h.Register(h.Path(), navBar, router)
}
