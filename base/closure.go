package main

import "fmt"

func main() {

	a := func() func() {
		//本曾的{}组成一个闭包代码块
		var i int = 0
		return func() {
			i++
			fmt.Println(i)
		}
	}() //要返回内层匿名函数

	a()
	a()
}
