package specification_test

import (
	"testing"
)

// Call      Argument type    Result
//
// len(s)    string type      string length in bytes
// [n]T, *[n]T      array length (== n)
// []T              slice length
// map[K]T          map length (number of defined keys)
// chan T           number of elements queued in channel buffer
//
// cap(s)    [n]T, *[n]T      array length (== n)
// []T              slice capacity
// chan T           channel buffer capacity

// 0 <= len(s) <= cap(s)

func TestLengthCap(t *testing.T) {
	stringWith4character := func() string { return "🐶中ßき" }()
	if len(stringWith4character) != 4 {
		t.Logf("stringWith4character actual length is %d(byte)", len(stringWith4character))
	}

	unbufferedChannel := make(chan int)
	t.Logf("unbuffered channel length %d, cap %d", len(unbufferedChannel), cap(unbufferedChannel))
	bufferedChannel := make(chan int, 10)
	t.Logf("buffered channel length %d, cap %d", len(bufferedChannel), cap(bufferedChannel))

	m := map[string]bool{
		"a": true,
		"b": true,
	}
	// map 没有cap操作
	t.Logf("map length %d", len(m))
	delete(m, "a")
	t.Logf("after delete, map length %d", len(m))

	s := make([]int, 2, 3)
	// 在cap范围内的
	if len(s) != 2 || cap(s) != 3 {
		t.Errorf("make([]int,2,3) failed")
	}

	prevSliceAddress := &s
	s = append(s, 4, 5, 6, 7, 8)
	nowSliceAddress := &s
	t.Logf("after append, slice length %d, cap %d, previous_address %p, now_address %p", len(s), cap(s),
		prevSliceAddress, nowSliceAddress)

}

// const (
// 	c1 = imag(2i)                    // imag(2i) = 2.0 is a constant
// 	c2 = len([10]float64{2})         // [10]float64{2} contains no function calls
// 	c3 = len([10]float64{c1})        // [10]float64{c1} contains no function calls
// 	c4 = len([10]float64{imag(2i)})  // imag(2i) is a constant and no function call is issued
// 	c5 = len([10]float64{imag(z)})   // invalid: imag(z) is a (non-constant) function call
// )
// var z complex128
