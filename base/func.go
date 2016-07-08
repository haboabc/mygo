package main

import "fmt"

/* 同类型不定长参数列表 */
func Myfunc1(args ...int) {
	for k, v := range(args) {
		fmt.Println(k, v)
	}
}

func Myfunc3(args ...int) {
	Myfunc1(args...) //不定长参数列表的引用方法
}

/* 不同类型，不定长参数列表 */
func Myfunc2(args ...interface{}) {
	for _, v := range(args) {
		switch v.(type) {
		case int:
			fmt.Println("it is a int value ", v)
		case string:
			fmt.Println("it is a string value ", v)
		default:
			fmt.Println("unknown type")
		}
	}
}

func main() {
	value1 := func(x, y int) int {
		var ret int
		ret = x + y
		return ret
	}(1, 2)

	fmt.Println(value1)

	value2 := func(x, y int) int {
		var ret int
		ret = x + y
		return ret
	}
	ret := value2(1, 2)
	fmt.Println(ret)

	Myfunc1(1, 2, 3, 4)
	Myfunc3(5, 6, 7, 8)

	Myfunc2("hello", 1, "world")
}
