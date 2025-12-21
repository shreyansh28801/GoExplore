package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("Its One")
	// 	//break -> here we don't need to write break, jsut like other langs
	// case 2:
	// 	fmt.Println("its two")
	// default:
	// 	fmt.Println("its none")
	// }

	// multiple cond. switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend, read books.")
	default:
		fmt.Println("its weekday, keep working")
	}

	//type switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("its type is int", t)
		case float32:
			fmt.Println("its type is float", t)
		case string:
			fmt.Println("its type is string", t)
		default:
			fmt.Println("its type id default", t)

		}
	}

	whoAmI("golang")
	whoAmI(12)
	whoAmI(12.8)
}
