package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Random Numbers")

	rand.Seed(time.Now().Unix())
	fmt.Print(rand.Intn(100), ",")
	rand.Seed(time.Now().UnixMilli())
	fmt.Print(rand.Intn(100))
	fmt.Println()

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Float64())

	rand.Seed(time.Now().UnixMicro())
	fmt.Print((rand.Float64()*5)+5, ",")
	rand.Seed(time.Now().UnixMilli())
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Println(r1.Intn(100))

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	fmt.Print(r2.Intn(100), ",")
	fmt.Println(r2.Intn(100))

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Println(r3.Intn(100))
}
