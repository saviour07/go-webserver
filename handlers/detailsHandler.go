package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saviour07/webserver/models"
)

// DetailsHandler structure
type DetailsHandler struct {
}

// RegisterHandler implements handlers.RegisterHandler
func (DetailsHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/details", detailsHandler)
}

func detailsHandler(writer http.ResponseWriter, request *http.Request) {
	namePage := template.Must(template.ParseFiles("./html/details.html"))

	if request.Method != http.MethodPost {
		namePage.Execute(writer, nil)
		return
	}

	vars := mux.Vars(request)
	name := vars["name"]

	fmt.Fprintf(writer, "Your name: %s\n", name)

	details := models.Details{
		Name: request.FormValue("name"),
		DOB:  request.FormValue("dob"),
	}

	fmt.Printf("Your details are\r\nName: %s\r\nDOB:%s", details.Name, details.DOB)

	namePage.Execute(writer, nil)
}
