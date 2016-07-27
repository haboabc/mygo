/* 实现栈，先进后出 */
package main

import "fmt"

type stack []int

func newStack() stack {
	return make(stack, 0)
}

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() (int){
	var index = len(*s) - 1
	if index < 0 {
		return 0
	}
	var ret = (*s)[0]
	*s = (*s)[1:]
	return ret
}


func main() {
	var stack = newStack()
	stack.push(1)
	stack.push(2)
	fmt.Println("NOW STACK STATUS:")
	fmt.Println("len :",len(stack))
	fmt.Println("cap :",cap(stack))
	for _, v := range(stack) {
		fmt.Println(v)
	}

	stack.pop()
	fmt.Println("len :",len(stack))
	fmt.Println("cap :",cap(stack))
	for _, v := range(stack) {
		fmt.Println(v)
	}
}
