package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------Tickers--------")

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				h, m, s := t.Clock()
				fmt.Println(
					"Tick at",
					fmt.Sprintf("%d:%d:%d:%d", h, m, s, t.Nanosecond()),
				)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
