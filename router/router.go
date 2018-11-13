package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saviour07/go-webserver/handlers"
)

var routerHandlers = []handlers.Handler{
	handlers.HomeHandler{},
	handlers.AboutHandler{},
}

// MuxRouter creates a new mux router and registers the handler functions
func MuxRouter() http.Handler {
	muxRouter := mux.NewRouter()

	for _, h := range routerHandlers {
		h.RegisterHandler(muxRouter)
	}

	return muxRouter
}
