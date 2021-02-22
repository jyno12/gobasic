package beginner_test

import (
	"fmt"
	"sync"
	"testing"
)

type field struct {
	name string
}

func (p *field) print(wg *sync.WaitGroup, idx int) {
	if wg != nil {
		defer wg.Done()
	}
	fmt.Println(p.name, idx)
}

func TestGoroutineFor(t *testing.T) {
	data := []*field{{"one"}, {"two"}, {"three"}}

	var wg sync.WaitGroup
	for idx, v := range data {
		wg.Add(1)
		go v.print(&wg, idx)
	}

	wg.Wait()
}
