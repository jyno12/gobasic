package specification_test

import (
	"sync"
	"testing"
)

// For a channel c, the built-in function close(c) records that no more values will be sent on the channel. It is an error if c is a receive-only channel. Sending to or closing a closed channel causes a run-time panic. Closing the nil channel also causes a run-time panic. After calling close, and after any previously sent values have been received, receive operations will return the zero value for the channel's type without blocking. The multi-valued receive operation returns a received value along with an indication of whether the channel is closed.

// 关闭只读的channel会报错
func TestCloseReceiveOnlyChannel(t *testing.T) {
	// 编译无法通过
	// ch := make(<-chan int)
	// if err := close(ch); err != nil {
	// 	t.Error(err)
	// }
}

// 往已经关闭的channel里发送数据会panic
func TestSendClosedChannel(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// close_channel_test.go:27: send on closed channel
			t.Log(r)
		}
	}()
	ch := make(chan int)
	close(ch)
	ch <- 1
}

// 关闭已经关闭的channel panic
func TestCloseClosedChannel(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// close_channel_test.go:40: close of closed channel
			t.Log(r)
		}
	}()
	ch := make(chan int)
	close(ch)
	close(ch)
}

// 接收已经关闭的channel，能一直接收到数据。数值是默认值，channel状态改变
func TestReceiveClosedChannel(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
	breakFor:
		for {
			select {
			case v, unclosed := <-ch:
				t.Logf("get value %d from channel, channel avaliable status is %v", v, unclosed)
				if !unclosed {
					v, unclosed = <-ch
					t.Logf("read from closed channel again, get value %d, avaliable status %v", v, unclosed)
					break breakFor
				}
			}
		}
	}()
	wg.Wait()
}
