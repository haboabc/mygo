package main

import "fmt"

var array [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	var slice1 []int
	var slice2 []int
	var slice3 []int
	var slice4 []int
	var slice5 []int

	slice1 = make([]int, 5, 10)
	slice2 = []int{0, 1, 2, 3}
	slice3 = make([]int, 5)
	slice4 = array[:5]
	slice5 = slice4[:5]

	slice1 = append(slice1, 1, 2, 3)
	slice2 = append(slice2, slice1...)

	copy(slice1, slice2)

	fmt.Println("slice1")
	fmt.Println("capicaty of slice1 is ", cap(slice1))
	fmt.Println("len of slice1 is ", len(slice1))
	for _, v := range slice1 {
		fmt.Println(v)
	}

	fmt.Println("slice2")
	fmt.Println("capicaty of slice1 is", cap(slice2))
	fmt.Println("len of slice2 is ", len(slice2))
	for _, v := range slice2 {
		fmt.Println(v)
	}

}
