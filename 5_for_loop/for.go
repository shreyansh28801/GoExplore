package main

import "fmt"

func main() {
	//classical for loop
	// for i := 0; i < 10; i++ {
	// 	if i == 4 {
	// 		continue
	// 	}
	// 	if i == 8 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	//impl. of while from 'for' loop
	// i := 1
	// for i<=3 {
	// 	fmt.Println(i+1)
	// 	i++
	// }

	//infinite loop
	// for {
	// 	fmt.Println("inf loop")
	// }


	//in go 1.22 range based looping is added
	for i:= range 3 {
		fmt.Println(i)
	}
}
