package main

import (
	"net/http"
	"text/template"
	"victor/golang/web/models"
)

var templates = template.Must(template.ParseGlob("web/store/templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request)  {
	products := models.SelectProducts()
	templates.ExecuteTemplate(w, "Index", products)
}
