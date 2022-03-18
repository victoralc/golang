package main

import "fmt"

var idade = 15

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Carlos Victor"
	fmt.Println(firstName)

	c := complex(3,4)
	r, im := real(c), imag(c)
	fmt.Println(r, im)

	//Scope of the package
	fmt.Println(idade)

	// method scope
	name, lastName := "Carlos Victor", "Alcantara Carleial"

	fmt.Println("First Name:", name)
	fmt.Println("Last Name:", lastName)

	//other type of declaration
	var (
		age int
		height float32
		nome string
	)

	age = 32
	height = 1.80
	nome = "Victor Alcantara"

	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(nome)

	// short declaration type
	color := 32
	width := 1.80
	label := "Victor Alcantara"

	fmt.Println(color)
	fmt.Println(width)
	fmt.Println(label)

	//Non initialized variables
	var developer, language string
	fmt.Println(developer, language)
}
