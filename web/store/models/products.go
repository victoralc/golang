package models

import (
	"log"
	db2 "victor/golang/web/store/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func SelectProducts() []Product {
	db := db2.ConnectDatabase()
	selectProducts, err := db.Query("select * from products order by id asc")
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

		p.Id = id
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
	db := db2.ConnectDatabase()
	insertInDb, err := db.Prepare("INSERT INTO products(name, description, quantity, price) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Println("Cannot create new product")
	} else {
		insertInDb.Exec(product.Name, product.Description, product.Quantity, product.Price)
		log.Println("Product created successfully", product)
	}

	defer db.Close()
}

func Delete(productId string) {
	db := db2.ConnectDatabase()
	deleteInDb, err := db.Prepare("DELETE FROM products WHERE ID = $1")
	if err != nil {
		panic(err.Error())
	} else {
		_, err := deleteInDb.Exec(productId)
		if err != nil {
			panic(err.Error())
		}
	}

	defer db.Close()
}

func FindById(productId string) Product {
	db := db2.ConnectDatabase()
	findProductByIdQuery, err := db.Query("select * from products where id = $1", productId)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	for findProductByIdQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = findProductByIdQuery.Scan(&id, &name, &description, &quantity, &price)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Quantity = quantity
		product.Price = price
	}

	defer db.Close()
	return product
}

func Update(productId int, name string, description string, price float64, quantity int) {
	db := db2.ConnectDatabase()
	updateProductStmt, err := db.Prepare("UPDATE products SET name =  $1, " +
		"description = $2, price = $3, quantity = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	} else {
		updateProductStmt.Exec(name, description, price, quantity, productId)
	}
	defer db.Close()
}
