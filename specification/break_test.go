package specification_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// break跳出for循环，之后的内容不会在执行
func TestBreakForLoop(t *testing.T) {
	for i := 1; i <= 10; i++ {
		if i > 2 {
			// [2,10]的内容不会输出
			break
		}
		t.Log(i)
	}
}

// break跳出for循环，跳到指定label
// 其中label必须包含跳出的for循环中
// 可以立即跳出多重循环，假设通过暴力遍历的方式寻找解
func TestBreakLabelForLoop(t *testing.T) {
	// var back bool
LoopLabel:
	for i := 1; i <= 10; i++ {
		t.Log(i)
		if i > 2 {
			// back = true
			break LoopLabel
		}
		t.Log(i)
	}
	t.Logf("loop label")

MultiLoopLabel:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i*i+j*j > 100 {
				break MultiLoopLabel
			}
			t.Log(i*10 + j)
		}
	}
	t.Logf("multi loop label")
	// LoopPrevLabel:
	// 	if back {
	// 		return
	// 	}
	t.Log("hello")
}

// break跳出switch_case
func TestBreakSwitchCase(t *testing.T) {
	zeroShowTimes := 0
	for zeroShowTimes <= 2 {
		randomNum := rand.Intn(2)
		switch randomNum {
		case 0:
			zeroShowTimes++
			t.Log("random number is 0")
			if zeroShowTimes > 2 {
				break
			}
			t.Log("random number is 0 repeated")

		case 1:
			t.Log("random number is 1")
		}
	}
}

// break跳出Switch_Case的场景
// Label在Switch_Case外面的场景
func TestBreakLabelSwitchCase(t *testing.T) {
	zeroShowTimes := 0
	// 由于只有一个label，可以维护一个退出状态
zeroShowTimesOverTwoLabel:
	for true {
		randomNum := rand.Intn(2)
		switch randomNum {
		case 0:
			subRandomNum := rand.Intn(2)
			switch subRandomNum {
			case 0:
				zeroShowTimes++
				t.Logf("double 0")
				if zeroShowTimes > 2 {
					break zeroShowTimesOverTwoLabel
				}
				t.Logf("double 0 repeated")
			case 1:
				t.Log("first 0 second 1")
			}

		case 1:
			t.Log("first loop random number 1")
		}
	}
	t.Log("final")
}

// break跳出Select语句
func TestBreakSelect(t *testing.T) {
	var even int
	var wg sync.WaitGroup
	randNumChan := make(chan int, 1)
	wg.Add(1)
	go func() {
		for even <= 1 {
			select {
			case num := <-randNumChan:
				t.Logf("num is %d", num)
				if num%2 == 0 {
					even++
				}
				if even > 1 {
					t.Logf("break")
					break
				}
				t.Logf("even count <= 2")
			default:

			}
		}
		// 剩余内容
		// _ = <-randNumChan
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for true {
			if even > 1 {
				wg.Done()
				return
			}
			num := rand.Intn(101)
			t.Logf("send num %d to channel", num)
			select {
			case randNumChan <- num:
			}
		}
	}()
	wg.Wait()
}

// break跳出Select语句
// 带Label语句的跳出
func TestBreakLabelSelect(t *testing.T) {
	t.Log("break跳出for循环")
BreakForLabel:
	for true {
		timeout := 1 * time.Second
		timer := time.After(timeout)
		select {
		case <-timer:
			nanoSecond := time.Now().Nanosecond()
			t.Logf("%v后，现在的纳秒时间为%v", timeout, nanoSecond)
			if nanoSecond%7 == 0 {
				break BreakForLabel
			}
		}
	}

	var breakCount int
	var normalOutCount int
	t.Log("break跳出select循环")
	for breakCount < 2 {
		timeout := time.Second * 1
		timer2s := time.After(timeout)
	BreakSelectLabel:
		select {
		case <-timer2s:
			normalOutCount++
			if time.Now().Unix()%7 == 0 {
				breakCount++
				t.Logf("第%d次通过break跳出，正常跳出次数增加但是不打印消息", breakCount)
				break BreakSelectLabel
			}
			t.Logf("第%d次正常跳出", normalOutCount)
		}
	}
	t.Log("break与select")
}
