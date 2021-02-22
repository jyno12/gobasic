package main

import (
	"fmt"
)

var (
	msg = `程序入口是main package开始，入口函数是main。` + "\n" +
		`package main必须包含main, 且传递任何的参数也不返回任何内容。` + "\n" +
		`注意：main函数是小写的、不可导出的。`
)

func main() {
	fmt.Println(msg)
}
