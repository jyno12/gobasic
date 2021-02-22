package beginner_test

import (
	"testing"
	"time"
)

type Escape struct {
	T *testing.T
}

//go:noinline
func (e Escape) returnPointerNestFunc() *int64 {
	ts := time.Now().UnixNano()
	e.T.Logf("函数内的局部变量原始的地址为：%v", &ts)
	wrapTs := ts
	e.T.Logf("函数内局部变量的的地址为：%v", &wrapTs)
	return &wrapTs
}

func (e Escape) ReturnPointer() {
	ts := e.returnPointerNestFunc()
	e.T.Logf("函数内的局部变量返回之后的指针对应的地址：%v", ts)
}

func TestEscape(t *testing.T) {
	// 发生逃逸的情况
	// 方法内局部变量指针返回
	// 发送指针或带有指针的值到channel中
	// 在一个切片上存储指针或带有指针的值
	// slice背后的数组重新分配
	// interface类型上调用方法
	e := Escape{T: t}
	e.ReturnPointer()
}
