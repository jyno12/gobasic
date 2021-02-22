package specification_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestReturn(t *testing.T) {
	// 正常返回，没有指定
	f1 := func() int {
		return rand.Intn(10)
	}
	t.Logf("random value: %d", f1())

	// 返回值命名
	f2 := func() (val int) {
		return 2
	}
	t.Logf("return value named: %d", f2())

	// 返回的时候是多值返回函数的返回
	f3 := func() (val int, nowTime string) {
		return func() (int, string) {
			return rand.Intn(10), fmt.Sprintf("%v", time.Now())
		}()
	}
	v, ts := f3()
	t.Logf("now %s return random value %d", ts, v)

	// 命名的返回值赋值
	f4 := func() (val int, nowTime string) {
		val = rand.Intn(10)
		nowTime = fmt.Sprintf("%v", time.Now())
		return
	}
	v4, ts4 := f4()
	t.Logf("now %s return random value %d", ts4, v4)
}

// Regardless of how they are declared, all the result values are initialized to the zero values
// for their type upon entry to the function. A "return" statement that specifies results sets the result parameters
// before any deferred functions are executed.
// 不管是否给返回值命名，当进入函数的时候返回值就已经创建并用0值初始化。
// return语句会在defer之前将返回值设置好。
func TestReturnDefer(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	{
		// 进入函数的时候，返回值就已经创建好
		f := func() int {
			expectReturnValue := rand.Intn(10)
			defer func() {
				t.Logf("expectReturnValue now is %d", expectReturnValue)
			}()

			defer func() {
				// defer函数在这里修改变量的值，所以局部变量修改了。但是函数返回值没有修改，函数返回值和变量是两个地址
				expectReturnValue = rand.Intn(100) + 100
			}()

			// 使用return语句的时候，return会将数据写入到函数返回值的变量地址上去
			return expectReturnValue
		}
		v := f()
		// 这里会返回的值是小于10的，说明返回值没有被修改
		t.Logf("call f return %d", v)
	}

	{
		// 和上面是相同的，只不过返回的值变成了某个变量的地址
		var variableAddress = new(*int)
		f := func(address **int) *int {
			expectReturnValue := new(int)
			*address = expectReturnValue
			t.Logf("expectReturnValue now address %p", expectReturnValue)

			defer func() {
				// 最后局部变量的值（变量地址）被修改
				t.Logf("after change expectReturnValue now address is %p", expectReturnValue)
			}()

			defer func() {
				// defer函数，这里修改变量值，即换个新的变量地址
				tmp := rand.Intn(100)
				expectReturnValue = &tmp
			}()

			// 使用return语句的时候，return会将数据写入到函数返回值的变量地址上去
			// 这里返回的是局部变量的地址
			return expectReturnValue
		}
		v := f(variableAddress)
		// 这里的值不过变成了指针而已
		t.Logf("call f return with variable address %p, local variable address %v", v, *variableAddress)
	}

	{
		// 命名的返回方式
		f := func() (r int) {
			expectReturnValue := rand.Intn(10)
			defer func() {
				t.Logf("expectReturnValue now is %d", expectReturnValue)
			}()

			defer func() {
				// defer函数在这里修改变量的值，所以局部变量修改了。但是函数返回值没有修改，函数返回值和变量是两个地址
				expectReturnValue = rand.Intn(100) + 100
			}()

			// 使用return语句的时候，return会将数据写入到函数返回值的变量地址上去
			return expectReturnValue
		}
		v := f()
		// 这里会返回的值是小于10的，说明返回值没有被修改
		t.Logf("call f return %d", v)
	}

	{ // 命名的返回方式
		f := func() (r int) {
			expectReturnValue := rand.Intn(10)
			defer func() {
				t.Logf("expectReturnValue now is %d", expectReturnValue)
			}()

			defer func() {
				// defer函数在这里修改变量的值，所以局部变量修改了。但是函数返回值没有修改，函数返回值和变量是两个地址
				expectReturnValue = rand.Intn(100) + 100
				// 这里将新的局部变量的值赋予到返回值上
				r = expectReturnValue
			}()

			// 使用return语句的时候，return会将数据写入到函数返回值的变量地址上去
			return expectReturnValue
		}
		v := f()
		// 这里返回的值是大于100，说明在defer函数中修改成功了
		t.Logf("call f return %d", v)
	}

	{
		// 命名的返回方式，在defer中修改命名的变量
		f := func() (r int) {
			expectReturnValue := rand.Intn(10)
			defer func() {
				t.Logf("expectReturnValue now is %d", expectReturnValue)
			}()

			defer func() {
				// defer函数在这里修改变量的值，所以局部变量修改了。但是函数返回值没有修改，函数返回值和变量是两个地址
				expectReturnValue = rand.Intn(100) + 100
				// 这里将新的局部变量的值赋予到返回值上
				r = expectReturnValue
			}()

			// 使用return语句的时候，return会将数据写入到函数返回值的变量地址上去
			return expectReturnValue
		}
		v := f()
		// 这里返回的值大于100说明修改成功了
		t.Logf("call f return %d", v)
	}
}
