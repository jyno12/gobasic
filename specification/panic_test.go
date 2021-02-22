package specification_test

import (
	"encoding/json"
	"runtime"
	"runtime/debug"
	"testing"
)

// panic之后recover

// 如果不在defer中调用recover，则recover返回值为nil
func TestPanicRecover(t *testing.T) {
	// if r := recover(); r == nil {
	// 	t.Logf("recover return expect nil, actual %v", r)
	// }
	// panic("cannot be recovered")
}

type CustomPanicStackInfo struct {
	StackInfo *CustomPanicStackInfo `json:"stack_info"`
	File      string                `json:"file"`
	Line      int                   `json:"line"`
	FuncName  string                `json:"caller_func"`
	Reason    *string               `json:"panic_reason"`
}

// recover success
// 通过defer recover来处理panic
// panic传递的信息被recover截取到
func TestPanicRecoverSuccess(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			b, _ := json.Marshal(r)
			t.Logf("%s", string(b))
		}
	}()

	PanicLayerG()
}

func PanicLayerG() {
	func() {
		func() {
			_, file, line, _ := runtime.Caller(0)
			panic(CustomPanicStackInfo{File: file, Line: line})
		}()
	}()
}
