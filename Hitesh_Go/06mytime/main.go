package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang.")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2000, time.September, 2, 2, 30, 15, 17, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
