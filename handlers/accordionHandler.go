package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saviour07/go-webserver/models"
)

// AccordionHandler structure
type AccordionHandler struct {
}

// Register implements handlers.Register()
func (AccordionHandler) Register(path string, navItems models.NavigationItems, router *mux.Router) {
	router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {

		navItems.Selected(path)

		pageData := models.Page{
			NavigationBar: navItems,
			Data: models.Accordion{
				Panels: []models.Panel{
					models.Panel{
						PanelName: "Options Section One",
						Options:   makeOptions(),
					},
					models.Panel{
						PanelName: "Options Section Two",
						Options:   makeOptions(),
					},
					models.Panel{
						PanelName: "Options Section Three",
						Options:   makeOptions(),
					},
				},
			},
		}

		html := template.Must(template.ParseFiles("./html/accordion.html", "./html/base/header.html", "./html/base/navbar.html", "./html/base/footer.html"))
		err := html.Execute(writer, pageData)
		if err != nil {
			fmt.Println(err)
		}
	})
}

func makeOptions() []models.Option {
	return []models.Option{
		models.Option{
			OptionName: "Option A",
		},
		models.Option{
			OptionName: "Option B",
		},
		models.Option{
			OptionName: "Option C",
		},
	}
}
