package simplemath

import (
	"testing" //for testing.T
)

func TestAdd1(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1, 2) failed")
	}
}
