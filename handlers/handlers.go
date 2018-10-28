package handlers

import (
	"github.com/gorilla/mux"
)

// Handler interface
type Handler interface {

	// RegisterHandler
	RegisterHandler(router *mux.Router)
}
