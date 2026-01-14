package main

import (
	"fmt"
	"slices"
)

func main() {
	//uninitialized slice is nill
	// var nums[]int
	// fmt.Println(nums==nil) //true
	// fmt.Println(len(nums)) //0

	// var nums = make([]int, 4, 5) //make method initializes array with initialize in 2nd arg. (here 4)
	// fmt.Println(nums==nil)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))
	// nums = append(nums, 1) //same as push_back in cpp
	// nums = append(nums, 2)
	// fmt.Println(cap(nums)) // it tell about the current cpaacity of nums basically that is reserved.


	//copy func for slices
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// fmt.Println(nums, nums2)
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)


	//slice operator
	// var nums = []int{1,2,3}
	// fmt.Println(nums[0:2]) //op -> [1 2]
	// fmt.Println(nums[:2]) //[1 2]
	// fmt.Println(nums[0:]) // [1 2 3]

	//slices package
	var nums1 = []int{1,2,3}
	var nums2 = []int{1,2,3}
	fmt.Println(slices.Equal(nums1, nums2)) //true

}