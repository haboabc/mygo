package main

import (
	"os"//for os.Args
	"fmt"//for fmt.Println
	"simplemath"//for simplemath.Add simplemath.Sqrt
	"strconv"//for strconv.Atoi
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\tThe commands are:\n\tadd\tAddition of two values.")
	fmt.Println("\tsqrt\tSquare root of a non-negative value.")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[1] {
		case "add":
			if len(args) != 4 {
				fmt.Println("USAGE: add <interger1> <interger2>")
				return
			}

			v1, err1 := strconv.Atoi(args[2])
			v2, err2 := strconv.Atoi(args[3])
			if err1 != nil || err2 != nil {
				fmt.Println("USAGE: args error...	add <interger1> <interger2>")
			}

			ret := simplemath.Add(v1, v2)
			fmt.Println("Result: ", ret)

		case "sqrt":
			if len(args) != 3 {
				fmt.Println("USAGE: sqlrt <interger>")
				return
			}

			v, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("Usage: sqrt <interger>")
				return
			}

			ret := simplemath.Sqrt(v)
			fmt.Println("Result: ", ret)

		default:
			Usage()
	}
}
