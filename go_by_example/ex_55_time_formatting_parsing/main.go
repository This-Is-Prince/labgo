package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	p := fmt.Println

	p("Time Formatting / Parsing")
	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, err := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	if err != nil {
		log.Fatal(err)
	}
	p(t1)

	p(t1.Format("3:04PM"))
	p(t1.Format("Mon Jan _2 15:04:05 2006"))
	p(t1.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, err := time.Parse(form, "8 41 PM")
	if err != nil {
		log.Fatal(err)
	}
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:05 2006"
	_, err = time.Parse(ansic, "8:41PM")
	p(err)
}
