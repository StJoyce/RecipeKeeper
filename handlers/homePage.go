package handlers

import (
	"html/template"
	"log"
	"net/http"
	"coding4women/data"
	
)
func HomePage(w http.ResponseWriter, r *http.Request){

	tmpl, err := template.ParseFiles("Frontend/homePage.html")
	if err != nil {
		log.Fatal(err, "ERROR: problem with home file path")
	}
	tmpl.Execute(w, data.AllRecipes)


	}

