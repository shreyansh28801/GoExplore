package main

import "fmt"

func main() {
	var name string = "golang"
	fmt.Println(name)

	//go automatically infers the type, if we initialize without type
	var lang = "golang"
	fmt.Println(lang)

	var age int = 30
	fmt.Println(age)

	//shorthand syntax
	shorthand_var := 31.5

	fmt.Println(shorthand_var)

	// variable declaration without intialization -> here we have to give type
	var temp string
	temp = "golang"
	fmt.Println(temp)

}
