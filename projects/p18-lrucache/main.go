package main

import (
	"fmt"

	"github.com/This-Is-Prince/lrucache/cache"
)

func main() {
	fmt.Println("Go LRU Cache")

	cache := cache.New()

	for _, word := range []string{"parrot", "avocado", "dragon fruit", "tree", "potato", "tomato", "tree"} {
		cache.Check(word)
		cache.Display()
	}
}
