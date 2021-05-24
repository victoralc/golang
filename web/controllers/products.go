package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"victor/golang/web/models"
)

var templates = template.Must(template.ParseGlob("web/store/templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request)  {
	products := models.SelectProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request)  {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request)  {
	product := models.Product{}
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceParsed, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Cannot convert price to float")
		}

		quantityParsed, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Cannot convert quantity to int")
		}

		product.Name = name
		product.Description = description
		product.Price = priceParsed
		product.Quantity = quantityParsed
	}
	models.Save(product)
	Index(w, r)
}