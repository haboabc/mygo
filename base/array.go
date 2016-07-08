package main

import "fmt"

func main() {
	var array1 [2]byte	//2元素数组， 类型为byte
	array1 = [2]byte{'c', 'd'}

	var array2 [2]struct{x, y int} = [2]struct{x, y int}{ {1, 2}, {3, 4} }

}
