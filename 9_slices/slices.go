package main

import (
	"fmt"
	"slices"
)

// Slices -> Dynamic Arrays
// MOST USED CONSTRUCT IN GO
// Useful methods
func main() {
	//uninitialized slice are nil by default
	var nums []int
	fmt.Println("Uninitialized array: ")
	fmt.Println(len(nums))

	//Way-1 to decalare slices
	var nums2 = make([]int, 0, 5)
	//capacity -> maximum number of elements which can fit
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)

	fmt.Println("Way-1: ")
	fmt.Println(nums2)
	fmt.Println(cap(nums2))
	fmt.Println(len(nums2))

	//Way-2 to decalare slices
	nums3 := []int{} //Key changes here
	nums3 = append(nums3, 1)
	nums3 = append(nums3, 2)

	fmt.Println("Way-2: ")
	fmt.Println(nums3)
	fmt.Println(cap(nums3))
	fmt.Println(len(nums3))

	//We can use index to add values in array as well
	nums2[0] = 4
	fmt.Println(nums2)

	//copy function
	var num = make([]int, 0, 5)
	num = append(num, 2)

	var num2 = make([]int, len(num))

	copy(num2, num)

	fmt.Println("Copy function: ")
	fmt.Println(num, num2)

	//Slice Operator
	var number = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice Operator: ")
	// from : to
	fmt.Println(number[0:2])
	fmt.Println(number[:1])
	fmt.Println(number[2:])

	//slice
	var numbers = []int{1, 2}
	var numbers2 = []int{1, 2}

	fmt.Println(slices.Equal(numbers, numbers2))

	//2-D slices
	var numm = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(numm)

}
