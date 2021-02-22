package beginner_test

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestStringLength(t *testing.T) {
	data := "é"
	fmt.Println(len(data))                    //prints: 3
	fmt.Println(utf8.RuneCountInString(data)) //prints: 2
}
