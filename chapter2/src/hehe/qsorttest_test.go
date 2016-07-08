package hehe
/* here is a bug, but i can't find it out */

import "testing"

func Testqsort0(t *testing.T) {
	values := []int{1, 3, 2, 5, 4}
	var ret int = 1

	Asort(values)

	for _, v := range(values) {
		if v != ret {
			t.Errorf("error")
		}
		ret++
   }
}

