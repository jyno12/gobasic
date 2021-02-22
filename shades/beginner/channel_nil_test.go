package beginner_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// nil的channel会阻塞
func TestNilChannelBlock(t *testing.T) {
	var ch chan int
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}

	//get first result
	fmt.Println("result:", <-ch)
	//do other work
	time.Sleep(2 * time.Second)
}

// select的动态变化
func TestNilChannel(t *testing.T) {
	inch := make(chan int)
	outch := make(chan int)

	// 这里不能使用make(chan bool)，没有缓冲区的channel会导致生产者生产之后就阻塞住，要消费者消费了之后才能继续
	// done := make(chan bool, 1)
	// consumeDone := make(chan bool, 1)
	done := make(chan bool, 1)
	consumeDone := make(chan bool, 1)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var in <-chan int = inch
		var out chan<- int
		var val int

	TestCaseOverLabel:
		for {
			select {
			// 步骤3：原来为nil的channel阻塞中，现在非空则可以执行
			// 设置in从nil->chan int
			case out <- val:
				if val == 2 {
					// 步骤5：结束测试用例
					t.Logf("最后一步")
					done <- true
					t.Log("本身结束信号发送完成")
					consumeDone <- true
					t.Log("结束信号发送完成")
				}

				out = nil
				in = inch

			case val = <-in:
				// 步骤2：in里面有数据了，现在设置in（chan int）为nil，然后阻塞住
				// 将out动态设置为非nil
				out = outch
				in = nil

			case <-done:
				// 步骤6：结束select
				break TestCaseOverLabel
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
	ConsumeForLabel:
		for {
			select {
			case <-consumeDone:
				t.Log("接收到结束信号，该结束了。")
				break ConsumeForLabel

			case r, ok := <-outch:
				if !ok {
					continue
				}
				fmt.Println("result:", r)
			default:
				break
			}
		}
		// for r := range outch {
		// 	fmt.Println("result:", r)
		// }

	}()

	time.Sleep(0)
	// 步骤1：往in里面写入数据
	inch <- 1
	// 步骤4：重新往in里面写入数据，并且跳回到步骤2，再执行步骤3
	inch <- 2

	// time.Sleep(3 * time.Second)
	wg.Wait()
}
