package specification_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// Execution of a "select" statement proceeds in several steps:
//
// 步骤1
// For all the cases in the statement, the channel operands of receive operations and the channel and
// right-hand-side expressions of send statements are evaluated exactly once, in source order,
// upon entering the "select" statement. The result is a set of channels to receive from or send to, and
// the corresponding values to send. Any side effects in that evaluation will occur irrespective of which (if any)
// communication operation is selected to proceed. Expressions on the left-hand side of a RecvStmt with
// a short variable declaration or assignment are not yet evaluated.
// 步骤2
// If one or more of the communications can proceed, a single one that can proceed is chosen via a
// uniform pseudo-random selection. Otherwise, if there is a default case, that case is chosen.
// If there is no default case, the "select" statement blocks until at least one of the communications can proceed.
// 步骤3
// Unless the selected case is the default case, the respective communication operation is executed.
// 步骤4
// If the selected case is a RecvStmt with a short variable declaration or an assignment,
// the left-hand side expressions are evaluated and the received value (or values) are assigned.
// 步骤5
// The statement list of the selected case is executed.

// select仅能针对channel操作，步骤
// 第一步：执行所有的case语句中有接收符号的channel（<-recvChannel）或者右值表达式（someVariable = <-recvChannel），仅仅执行一次。
// 这一步按照代码顺序偶从上到下的定义顺序。
// 结果就是从channel中获取到或者发送数据的数据，
// 第二步：如果一个或多个通信可以被执行，通过统一伪随机算法选取一个可执行case；否则运行默认的case。
// 如果没有default case，则select会阻塞直到有一个能执行。
// 第三步：除非选择的case分支是default，否则通信过程会被执行。因为其他的case分支一定是和channel相关，就一定存在通信。
// 这里说的通信过程就是从channel中发送或接收信息。
// 第四步：如果选择的case分支是一个伴随着变量变量定义的语句（shorVariable:=<-recvChannel）或者是一个声明语句，
// 左边的表达式会在接收时候执行（shortVariable=<-recvChannel）或者赋值。
// 第五步：被选择case才正式执行。

func TestSelectChannelSend(t *testing.T) {
	sendChan := make(chan int)
	sendChan1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		timer := time.NewTimer(20 * time.Millisecond)
	forBreak:
		for {
			val := rand.Intn(100)
			select {
			// 发送数据的channel
			case sendChan <- val:
				t.Logf("produce data %d", val)
				time.Sleep(5 * time.Millisecond)
			// 	发送数据
			case sendChan1 <- val:
				t.Logf("produce1 data %d", val)
			// 这个没有左边的赋值
			case <-timer.C:
				t.Logf("send %v timeout", 20*time.Millisecond)
				sendChan <- val
				t.Logf("last send %v", val)
				close(sendChan)
				close(sendChan1)
				break forBreak
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		var countFromChan1 int
	forBreak:
		for {
			select {
			// 接收数据的channel
			case val, ok := <-sendChan:
				t.Logf("consume data %d", val)
				time.Sleep(3 * time.Millisecond)
				if !ok {
					t.Logf("last consume data %d", val)
					break forBreak
				}

			// 左边的只能是：赋值、声明和省略左边
			case countFromChan1 = <-sendChan1:
				t.Logf("consume1 data")

			// 	不满足上面的条件就执行这个
			default:
				t.Logf("default case, sleep for 3 millis second")
				time.Sleep(3 * time.Millisecond)
				continue
			}
		}
	}()
	wg.Wait()
}
