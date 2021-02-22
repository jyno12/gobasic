package beginner_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonEndNewLine(t *testing.T) {
	data := map[string]int{"key": 1}

	// encode会自带换行
	var b bytes.Buffer
	json.NewEncoder(&b).Encode(data)

	raw, _ := json.Marshal(data)

	if b.String() == string(raw) {
		fmt.Println("same encoded data")
	} else {
		fmt.Printf("'%s' != '%s'\n", raw, b.String())
		//prints:
		//'{"key":1}' != '{"key":1}\n'
	}
}
