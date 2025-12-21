package main

import "fmt"

var name = "golang"
const age = 30

// shorthand_var := 67 -> not allowed, shorthand only inside funciton

func main() {
	const name1 = "golang"
	// name = "java" //not allowed

	//constant grouping -> if multiple constants need to be intialize at same time
	const (
		port = 3000
		host = "localhost"
	)
	// port = 3001 -> not allowed bcz its constant

	fmt.Println(port, host)
}