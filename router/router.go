package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saviour07/go-webserver/controller"
	"github.com/saviour07/go-webserver/models"
)

var controllers = []controller.Controller{
	&controller.HomeController{},
	&controller.AboutController{},
	&controller.AccordionController{},
}

// MuxRouter creates a new mux router and registers the handler functions
func MuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	registerHandlers(muxRouter)
	return muxRouter
}

func registerHandlers(muxRouter *mux.Router) {
	navBar := buildNavBar()
	for idx := range controllers {
		controllers[idx].RegisterHandler(navBar, muxRouter)
	}
}

func buildNavBar() models.NavigationItems {
	var navBar models.NavigationItems
	for idx := range controllers {
		navBar.NavItems = append(navBar.NavItems, controllers[idx].NavItem())
	}
	return navBar
}
