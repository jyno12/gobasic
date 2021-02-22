package specification_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	// 普通定义，默认为nil
	if func() chan int { var nilChan chan int; return nilChan }() == nil {
		t.Log("定义chan int未初始化，为nil")
	}

	// 初始化chan大小，未指定大小或者位0，则channel没有缓存空间，必须理解消费与生产
	defaultChan := make(chan int)
	t.Logf("默认的channel大小为：%d，长度：%d", cap(defaultChan), len(defaultChan))
	// 消费者
	go func() {
		select {
		case num := <-defaultChan:
			t.Logf("消费channel中的数据：%d", num)
		}
	}()
	// 生产者
	defaultChan <- rand.Intn(101)
}

func TestBufferChannel(t *testing.T) {
	// 初始化channel的大小
	bufferChannel := make(chan int, 3)
	consumeOverChannel := make(chan bool)
	t.Logf("默认的channel大小为：%d，长度：%d", cap(bufferChannel), len(bufferChannel))

	var wg sync.WaitGroup
	wg.Add(1)
	// 消费者
	go func() {
	ConsumeOverLabel:
		for true {
			select {
			case num := <-bufferChannel:
				t.Logf("消费缓冲channel中的数据：%d", num)
			case <-consumeOverChannel:
				t.Logf("生产者生产完了")
				break ConsumeOverLabel
			}
		}
		wg.Done()
	}()

	// 生产者
	rand.Seed(time.Now().UnixNano())
	bufferChannel <- rand.Intn(10)
	bufferChannel <- rand.Intn(101)
	bufferChannel <- rand.Intn(70001)

	wg.Add(1)
	go func() {
		for true {
			time.Sleep(200 * time.Millisecond)
			if time.Now().UnixNano()%17 == 0 {
				consumeOverChannel <- true
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

func TestSingleDirectionChannel(t *testing.T) {
	// 定义单向的channel
	var sendOnlyChan chan<- int
	var receiveOnlyChan <-chan int
	// 设置一个缓冲大小为1的channel
	bufferChannel := make(chan int, 1)

	sendOnlyChan = func() chan<- int {
		// 对外只提供写功能
		return bufferChannel
	}()

	receiveOnlyChan = func() <-chan int {
		// 对外只提供读功能
		return bufferChannel
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	// 消费者
	go func() {
		// 通过只读的channel，避免写操作
		select {
		case num := <-receiveOnlyChan:
			t.Logf("从只读的channel中读取到：%d", num)
		}
		wg.Done()
	}()

	// 生产者
	rand.Seed(time.Now().UnixNano())
	sendOnlyChan <- rand.Intn(101)
	wg.Wait()
}
