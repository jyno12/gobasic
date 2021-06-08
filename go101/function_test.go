package go101_test

import (
	"fmt"
	"testing"
)

// https://twitter.com/go100and1/status/1364210806687027204

// 函数返回值
func TestFunctionCall(t *testing.T) {
	fmt.Println(f(1))  // 0,1
	fmt.Println(f1(1)) // 0,0
}

// 单个下划线是忽略，所以返回的会是0值；双下划线是正常的变量
func f(x int) (_, __ int) {
	_, __ = x, x
	return
}

// 单个下划线是忽略
func f1(x int) (_, _ int) {
	_, _ = x, x
	return
}
