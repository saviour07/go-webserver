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

// Register implements handlers.Register()
func (HomeHandler) Register(path string, navItems models.NavigationItems, router *mux.Router) {
	router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {

		navItems.Selected(path)

		pageData := models.Page{
			NavigationBar: navItems,
			Data: models.Home{
				Message: "Hello, World!",
			},
		}

		html := template.Must(template.ParseFiles("./html/home.html", "./html/base/header.html", "./html/base/navbar.html", "./html/base/footer.html"))
		err := html.Execute(writer, pageData)
		if err != nil {
			fmt.Println(err)
		}
	})
}
