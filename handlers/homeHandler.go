package handlers

import (
	"fmt"
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
	html := template.Must(template.ParseFiles("./html/home.html", "./html/base/header.html", "./html/base/navbar.html", "./html/base/footer.html"))
	homePage := models.Home{
		Message: "Hello, World!",
	}
	err := html.Execute(writer, homePage)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
