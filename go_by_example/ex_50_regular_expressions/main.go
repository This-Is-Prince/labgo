package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
)

func main() {
	fmt.Println("Regular Expressions")

	match, err := regexp.MatchString("p([a-z]+)ch", "peach")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(match)

	r, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
