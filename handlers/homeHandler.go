package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saviour07/go-webserver/models"
)

// HomeHandler structure
type HomeHandler struct {
}

// RegisterHandler implements handlers.RegisterHandler()
func (HomeHandler) RegisterHandler(path string, navItems []models.NavItem, router *mux.Router) {
	router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {

		var idx int
		for idx = range navItems {
			navItems[idx].Selected = false
			if navItems[idx].Path == path {
				navItems[idx].Selected = true
			}
		}

		pageData := models.Page{
			NavigationBar: navItems,
			Data: models.Home{
				Message: "Hello, World!",
			},
		}

		html := template.Must(template.ParseFiles("./html/home.html", "./html/base/header.html", "./html/base/navbar.html", "./html/base/footer.html"))
		err := html.Execute(writer, pageData)
		if err != nil {
			fmt.Printf(err.Error())
		}
	})
}
