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

// Register implements handlers.Register()
func (AboutHandler) Register(path string, navItems models.NavigationItems, router *mux.Router) {
	router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {

		navItems.Selected(path)

		pageData := models.Page{
			NavigationBar: navItems,
			Data: models.About{
				AboutDetail: "About Details",
			},
		}

		html := template.Must(template.ParseFiles("./html/about.html", "./html/base/header.html", "./html/base/navbar.html", "./html/base/footer.html"))
		err := html.Execute(writer, pageData)
		if err != nil {
			fmt.Println(err)
		}
	})
}
