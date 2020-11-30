package intergers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Adder(3, 5)
	expected := 8

	if sum != expected {
		t.Errorf("Expected %d but got %d", expected, sum)
	}
}

// 例示
func ExampleAdder() {
	sum := Adder(2, 2)
	fmt.Println(sum)
	// Output: 4
}