package main

import (
	"net/http"
	"github.com/Creator/Medinfo/handlers"
)

func main() {

	http.HandleFunc("/login", handlers.LoginHandler)
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":1337", nil)

}
