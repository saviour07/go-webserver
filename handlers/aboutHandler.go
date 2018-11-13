package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saviour07/go-webserver/models"
)

// AboutHandler structure
type AboutHandler struct {
}

// RegisterHandler implements handlers.RegisterHandler
func (AboutHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/about", aboutHandler)
}

func aboutHandler(writer http.ResponseWriter, request *http.Request) {
	html := template.Must(template.ParseFiles("./html/about.html", "./html/base/header.html", "./html/base/navbar.html", "./html/base/footer.html"))
	aboutPage := models.About{
		AboutDetail: "about details",
	}
	err := html.Execute(writer, aboutPage)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
