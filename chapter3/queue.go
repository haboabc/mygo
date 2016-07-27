package main

import "fmt"

type queue []int

func newQueue() queue {
	return make(queue, 0)
}

func (q *queue) In(x int) {
	*q = append(*q, x)
}

func (q *queue) Out() (ret int) {
	if len(*q) <= 0 {
		return 0
	}

	ret = (*q)[0]
	*q = (*q)[1:]

	return
}

func main() {
	var queue = newQueue()

	fmt.Println("NOW QUEUE STATUS:")
	fmt.Println("len :", len(queue))
	fmt.Println("cap :", cap(queue))
	for _, v := range(queue) {
		fmt.Println(v)
	}

	queue.In(1)
	queue.In(2)
	queue.In(3)
	fmt.Println("NOW QUEUE STATUS:")
	fmt.Println("len :", len(queue))
	fmt.Println("cap :", cap(queue))
	for _, v := range(queue) {
		fmt.Println(v)
	}

	queue.Out()
	fmt.Println("NOW QUEUE STATUS:")
	fmt.Println("len :", len(queue))
	fmt.Println("cap :", cap(queue))
	for _, v := range(queue) {
		fmt.Println(v)
	}
}
