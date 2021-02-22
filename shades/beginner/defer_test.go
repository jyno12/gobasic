package beginner_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDeferCallArgumentEvaluation(t *testing.T) {
	{
		var i = 1

		// defer将要执行的内容已经确定
		defer fmt.Println("result => ", func() int { return i * 2 }())
		i++
	}

	{
		i := 1
		// 要执行的内容已经确定，确定的是指针
		defer func(in *int) { fmt.Println("result =>", *in) }(&i)

		i = 2
	}
}

func TestDeferCallExecution(t *testing.T) {
	if len(os.Args) != 2 {
		os.Exit(-1)
	}

	start, err := os.Stat(os.Args[1])
	if err != nil || !start.IsDir() {
		os.Exit(-1)
	}

	var targets []string
	filepath.Walk(os.Args[1], func(fpath string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		targets = append(targets, fpath)
		return nil
	})

	{
		// 方式一、这里调用defer会提示问题，因为在for循环中导致失效
		for _, target := range targets {
			f, err := os.Open(target)
			if err != nil {
				fmt.Println("bad target:", target, "error:", err) //prints error: too many open files
				break
			}
			defer f.Close() //will not be closed at the end of this code block
			//do something with the file...
		}
	}

	{
		// 方式二、这里通过外包一个函数来处理
		for _, target := range targets {
			// 放在一个函数形式中，这样就以调用defer
			func() {
				f, err := os.Open(target)
				if err != nil {
					fmt.Println("bad target:", target, "error:", err)
					return
				}
				defer f.Close() //ok
				//do something with the file...
			}()
		}
	}
}
