package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saviour07/go-webserver/controller"
	"github.com/saviour07/go-webserver/models"
)

var controllers = []controller.Controller{
	controller.HomeController{},
	controller.AboutController{},
}

// MuxRouter creates a new mux router and registers the handler functions
func MuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	registerHandlers(muxRouter)
	return muxRouter
}

func registerHandlers(muxRouter *mux.Router) {
	navBar := buildNavBar()
	for _, c := range controllers {
		c.RegisterHandler(navBar, muxRouter)
	}
}

func buildNavBar() []models.NavItem {
	var navBar []models.NavItem
	for _, c := range controllers {
		navBar = append(navBar, c.NavItem())
	}
	return navBar
}
