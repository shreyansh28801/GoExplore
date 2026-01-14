package main

import "fmt"

//numbered sequence of specific length
func main() {
	// var nums [4]int //default value in int is zero
	// //arrray length
	// fmt.Println(len(nums))
	// nums[0] = 1
	// fmt.Println(nums[0]) //get by index
	// fmt.Println(nums) // print array

	// var arr[3] bool //default value is false
	// fmt.Println(arr)

	// var vals [5]string //default value is empty string
	// fmt.Println(vals)

	//to declare array in single line
	// nums := [3]int{1,2,3}
	// fmt.Println(nums)

	//to declare array in 2d space
	nums := [3][2]int{{1,2},{3,4},{5,6}}
	fmt.Println(nums)

	//fixed sized array - the above mentod of declaring array is used when we know the size of array in advance
}