package beginner_test

import (
	"fmt"
	"testing"
)

type data struct {
}

// If you have zero-sized variables they might share the exact same address in memory.
func TestZeroSizeVariable(t *testing.T) {
	a := &data{}
	b := &data{}

	if a == b {
		fmt.Printf("same address - a=%p b=%p\n", a, b)
		//prints: same address - a=0x1953e4 b=0x1953e4
	}
}
