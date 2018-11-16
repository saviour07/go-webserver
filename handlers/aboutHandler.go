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
	navItems []models.NavItem
}

// AddNavBar implements handlers.AddNavBar()
func (a *AboutHandler) AddNavBar(navItems []models.NavItem) {
	a.navItems = navItems
}

// RegisterHandler implements handlers.RegisterHandler
func (a AboutHandler) RegisterHandler(path string, router *mux.Router) {
	router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {

		pageData := models.Page{
			NavigationBar: a.navItems,
			Data: models.About{
				AboutDetail: "About Details",
			},
		}

		html := template.Must(template.ParseFiles("./html/about.html", "./html/base/header.html", "./html/base/navbar.html", "./html/base/footer.html"))
		err := html.Execute(writer, pageData)
		if err != nil {
			fmt.Printf(err.Error())
		}
	})
}
