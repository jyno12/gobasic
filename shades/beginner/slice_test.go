package beginner_test

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSliceCorruption(t *testing.T) {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB (not ok)

	fmt.Println("new path =>", string(path))

	// 符合预期
	{
		path := []byte("AAAA/BBBBBBBBB")
		sepIndex := bytes.IndexByte(path, '/')
		dir1 := path[:sepIndex:sepIndex] //full slice expression
		// 最后一个参数是slice的大小，如果再添加内容则会重新分配空间，因此不会冲突
		dir2 := path[sepIndex+1:]
		fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
		fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

		dir1 = append(dir1, "suffix"...)
		path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})

		fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
		fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB (ok now)

		fmt.Println("new path =>", string(path))
	}
}
