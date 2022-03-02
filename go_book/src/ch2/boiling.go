package ch2

import "fmt"

const boilingF = 212.0

// Boiling prints the boiling point of water.
func Boiling() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}
