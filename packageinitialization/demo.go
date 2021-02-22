package packageinitialization

import (
	"fmt"
	"math/rand"
)

var (
	a = c + b // == 9
	b = f()   // == 4
	c = f()   // == 5
	d = 3     // == 5 after initialization has finished
)

func f() int {
	d++
	return d
}

var (
	// 最后初始化
	x = final()
	// 按照定义顺序初始化
	ax, bx = f1()
	// 按照定义顺序初始化
	cx = f2()

	dx = f3()
)

// init可以定义多个，按照定义的顺序执行
// 所以Go程序中函数是可以重命名定义的
func init() {
	fmt.Println("init -1")
}

func init() {
	fmt.Printf("init 0 step 0, init function\n")
}

func init() {
	fmt.Println("init 1")
}

func final() int {
	valueX := ax + bx + cx
	fmt.Printf("step 3, initialize x by a, b c sum = %d\n", valueX)
	return valueX
}

func f1() (int, int) {
	valueA, valueB := rand.Intn(10), rand.Intn(10)+10
	valueB += dx
	fmt.Printf("step 1 initialize a, b with random value(ax: %v, bx: %v)\n", valueA, valueB)
	return valueA, valueB
}

func f2() int {
	valueC := rand.Intn(100) + 100
	fmt.Printf("step 2 initialize c with random value(cx: %v)\n", valueC)
	return valueC
}

func f3() int {
	valueD := rand.Intn(100) + 100
	fmt.Printf("step 2 initialize c with random value(dx: %v)\n", valueD)
	return valueD
}
