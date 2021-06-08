package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ParentTimeoutCtx() {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	timer := time.NewTimer(time.Second * 3)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go sonTimeCtx(timeoutCtx, wg)

	select {
	// 如果结束
	case <-timeoutCtx.Done():
		fmt.Println("timeout")

	case <-timer.C:
		fmt.Println("timer")
	}
	wg.Wait()
	// 睡一会，看看sonTimeCtx是不是还在继续工作
	time.Sleep(time.Minute)
}

func sonTimeCtx(ctx context.Context, wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("close son time context function")
		wg.Done()
	}()
	resultChannel := make(chan int)
	go func() {
		now := time.Now()
		var prevSecond int
		for time.Since(now) < time.Second*5 {
			if currentSecond := time.Now().Second(); currentSecond != prevSecond {
				prevSecond = currentSecond
				fmt.Printf("%v\n", prevSecond)
			}
		}
		resultChannel <- 1
	}()

	select {
	case <-ctx.Done():
		fmt.Println("context done")
		return
	case <-resultChannel:
		fmt.Println("get result")
		return
	}
}
