package handlers

import (
	//"fmt"
	"html/template"
	"net/http"

	"github.com/Creator/Medinfo/static/goget"
)

var data *goget.DataStruct

//IndexHandler Index page handler function
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	tmpl.Execute(w, data)

}

func blank() (data *goget.DataStruct) {

	image := "static/images/Gopher Google.jpg"
	name := "Gopher"
	surname := "Google"
	middlename := "Go"
	info := "The one who does the work is the Gopher"

	data = &goget.DataStruct{

		Image:       image,
		Name:        name,
		Surname:     surname,
		Middlename:  middlename,
		Both:        name + " " + surname,
		Information: info,
	}

	return data

}
