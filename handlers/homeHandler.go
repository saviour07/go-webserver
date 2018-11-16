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
	navItems []models.NavItem
}

// AddNavBar implements handlers.AddNavBar()
func (h *HomeHandler) AddNavBar(navItems []models.NavItem) {
	h.navItems = navItems
}

// RegisterHandler implements handlers.RegisterHandler
func (h HomeHandler) RegisterHandler(path string, router *mux.Router) {
	router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {

		pageData := models.Page{
			NavigationBar: h.navItems,
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
