package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// VariableHandler structure
type VariableHandler struct {
}

// RegisterHandler implements handlers.RegisterHandler
func (VariableHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/{var}", variableHandler)
}

func variableHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	v := vars["var"]
	fmt.Fprintf(writer, "Demonstrate reading a variable from the URL: %s", v)
}
