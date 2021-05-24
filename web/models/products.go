package models

import (
	"victor/golang/web/db"
)

type Product struct {
	Name string
	Description string
	Price float64
	Quantity int
}

func SelectProducts() []Product {
	db := db.ConnectDatabase()
	selectProducts, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &quantity, &price)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Quantity = quantity
		p.Price = price

		products = append(products, p)
	}
	defer db.Close()
	return products
}