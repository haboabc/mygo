package main

import "fmt"

func main() {
	var v1 int	//整形
	var v2 string	//字符串
	var v3 [10]int	//10元素数组
	var v4 []int //切片
	var v5 map[string]int //键值对
	var v6 struct { //结构
		f int
	}
	var v7 func(i, j int) int //函数
	var v8 *int //指针


	var (
		v9 int
		v10 string
	)

	var v11 int = 10
	var v12 = 10
	v13 := 10

	const Pi float64 = 3.1415926
	const Qi = 3.1415926

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
}
