package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	models2 "victor/golang/web/store/models"
)

var templates = template.Must(template.ParseGlob("web/store/templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request)  {
	products := models2.SelectProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request)  {
	templates.ExecuteTemplate(w, "New", nil)
}

func Delete(w http.ResponseWriter, r *http.Request)  {
	productId := r.URL.Query().Get("id")
	models2.Delete(productId)
	templates.ExecuteTemplate(w, "Index", models2.SelectProducts())
}

func Edit(w http.ResponseWriter, r *http.Request)  {
	productId := r.URL.Query().Get("id")
	product := models2.FindById(productId)
	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConverted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Cannot convert id into int", err)
		}

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Cannot convert price to float64", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Cannot convert quantity into int", err)
		}

		models2.Update(idConverted, name, description, priceConverted, quantityConverted)

	}
	templates.ExecuteTemplate(w, "Index", models2.SelectProducts())
}

func Insert(w http.ResponseWriter, r *http.Request)  {
	product := models2.Product{}
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
	models2.Save(product)
	Index(w, r)
}