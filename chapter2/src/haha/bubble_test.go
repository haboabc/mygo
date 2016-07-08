package haha

import "testing"

func TestBubble0(t *testing.T) {
	values := []int{1, 3, 2, 5, 4}
	var ret int = 5

	Bubble(values)

	for _, v := range(values) {
		if v != ret {
			t.Errorf("err")
		}

		ret--
	}
}
