package main

import "fmt"

func main() {
	var map1 map[string]int = map[string]int{
		"1": 1,
		"2": 2,
	}

	var map2 map[string]int = make(map[string]int, 10)
	var map3 = make(map[string]int)

	map2["1"] = 1
	map2["2"] = 2

	delete(map2, "1")

	intval, ok := map2["1"]
	if ok {
		fmt.Println("find it ", intval)
	} else {
		fmt.Println("not find")
	}
}
