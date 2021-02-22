package specification_test

import (
	"sync"
	"testing"
)

func TestForRange(t *testing.T) {
	ForRangeCase1(t)
	ForRangeCase2(t)
}

func ForRangeCase1(t *testing.T) {
	var wg sync.WaitGroup
	// index:2, value:0, render_index:2, render_value:0
	// index:1, value:1, render_index:2, render_value:0
	// index:0, value:2, render_index:2, render_value:0
	for k, v := range []int{2, 1, 0} {
		wg.Add(1)
		go func(index, value int) {
			defer wg.Done()
			t.Logf("index:%d, value:%d, render_index:%d, render_value:%d", index, value, k, v)
		}(k, v)
	}
	wg.Wait()
	t.Logf("-----")
}

// 以闭包的形式在for循环中访问，会产生未定义的问题。
// 协程读取到内容不能够确定
func ForRangeCase2(t *testing.T) {
	var wg sync.WaitGroup
	arrCount := 10000
	arr := make([]*int, arrCount)
	for i := 0; i < arrCount; i++ {
		i := i
		arr[i] = &i
	}

	type Pair struct {
		Index int
		Value int
	}

	totalChan := make(chan int)
	statChan := make(chan Pair)
	for k, v := range arr {
		wg.Add(1)
		go func(index int, value *int) {
			defer wg.Done()
			defer func() { totalChan <- 1 }()
			if k != arrCount-1 {
				// t.Logf("index:%d, value:%d, render_index:%d, render_value:%d", index, *value, k, *v)
				statChan <- Pair{Index: k, Value: *v}
			}
		}(k, v)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		statMap := make(map[int]Pair)
		var total int
	breakFor:
		for {
			select {
			case p := <-statChan:
				statMap[p.Index] = p
			case <-totalChan:
				total++
				if total == arrCount {
					t.Logf("总共执行的次数：%v，len(index_map):%d", total, len(statMap))
					break breakFor
				}
			}
		}

		var diff int
		for _, stat := range statMap {
			if stat.Index != stat.Value {
				diff++
				// t.Logf("(index:%v, value:%v)", stat.Index, stat.Value)
			}
		}
		t.Logf("diff %d", diff)
	}()
	wg.Wait()
}
