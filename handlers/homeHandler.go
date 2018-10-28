package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saviour07/webserver/models"
)

// HomeHandler structure
type HomeHandler struct {
}

// RegisterHandler implements handlers.RegisterHandler
func (HomeHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/", homeHandler)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	html := template.Must(template.ParseFiles("./html/home.html"))
	homePage := models.Home{
		PageTitle: "You're home!",
	}
	html.Execute(writer, homePage)
}
