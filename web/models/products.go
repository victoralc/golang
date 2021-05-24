package models

import (
	"log"
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

func Save(product Product) {
	db := db.ConnectDatabase()
	insertInDb, err := db.Prepare("INSERT INTO products(name, description, quantity, price) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Println("Cannot create new product")
	} else {
		insertInDb.Exec(product.Name, product.Description, product.Quantity, product.Price)
		log.Println("Product created successfully", product)
	}

	defer db.Close()
}