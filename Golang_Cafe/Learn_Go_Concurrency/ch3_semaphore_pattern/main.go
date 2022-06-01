package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

type Task struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	var t Task
	start := time.Now()
	var totalRequest = 100
	wg := &sync.WaitGroup{}
	errGroup, _ := errgroup.WithContext(context.Background())
	wg.Add(totalRequest)
	// sem := make(chan bool, 10)
	sem := semaphore.NewWeighted(10)
	for i := 1; i <= totalRequest; i++ {
		fmt.Println(runtime.NumGoroutine())
		// sem <- true
		if err := sem.Acquire(context.Background(), 1); err != nil {
			log.Fatal(err)
		}
		/*
		   go func(i int) {
		   			defer wg.Done()
		   			// defer func() {
		   			// 	<-sem
		   			// }()
		   			defer sem.Release(1)
		   			if res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i)); err != nil {
		   				log.Fatal(err)
		   			} else {
		   				defer res.Body.Close()
		   				if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		   					log.Fatal(err)
		   				}
		   				fmt.Println(i, t.Title)
		   			}
		   		}(i)
		*/
		i := i
		errGroup.Go(func() error {
			defer wg.Done()
			// defer func() {
			// 	<-sem
			// }()
			defer sem.Release(1)
			if res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i)); err != nil {
				return err
			} else {
				defer res.Body.Close()
				if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
					return err
				}
				fmt.Println(i, t.Title)
			}
			return nil
		})
	}
	if err := errGroup.Wait(); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
	fmt.Println("total time: ", time.Since(start))
}
