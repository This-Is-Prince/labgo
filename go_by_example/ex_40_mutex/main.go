package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mut      sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.counters[name]++
}

func main() {
	fmt.Println("-------Mutex-------")

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(6)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	go doIncrement("b", 10000)
	go doIncrement("b", 10000)
	go doIncrement("c", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
