package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("------------Strings And Runes------------")

	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\n Using DecodeRuneInString")
	examineRune := func(r rune) {
		if r == 't' {
			fmt.Println("Found tee")
		} else if r == 'ส' {
			fmt.Println("Found so sua")
		}
	}
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}
