package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type cache struct {
	mut *sync.RWMutex
	m   map[string]int
}

var inMemoryCache = &cache{m: make(map[string]int), mut: &sync.RWMutex{}}

func main() {
	fmt.Println("Tutorial 4 - Mutex")

	wg := &sync.WaitGroup{}
	rand.Seed(time.Now().UnixNano())
	wg.Add(100)
	for i := 0; i < 100; i++ {
		// saveSlow(fmt.Sprintf("user_id_%d", rand.Intn(100)), rand.Intn(100))
		go save(fmt.Sprintf("user_id_%d", rand.Intn(100)), rand.Intn(100), wg)
	}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			inMemoryCache.mut.RLock()
			fmt.Println(inMemoryCache.m)
			inMemoryCache.mut.RUnlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(inMemoryCache.m)
}

func save(k string, v int, wg *sync.WaitGroup) {
	inMemoryCache.mut.Lock()
	inMemoryCache.m[k] = v
	inMemoryCache.mut.Unlock()
	wg.Done()
}

func saveSlow(k string, v int) {
	inMemoryCache.m[k] = v
}
