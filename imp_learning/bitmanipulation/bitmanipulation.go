package bitmanipulation

import (
	"fmt"
	"math"
)

/*
Instead of thinking of a number like 12 as just "twelve," you think of it as its binary representation: 1100. Working at this level is useful for:

Flags: Packing multiple boolean (on/off) states into a single integer.

Performance: Bitwise operations are often faster than arithmetic operations like multiplication or division.

Low-level Control: Interacting with hardware, protocols, or binary file formats.
*/

// 9 -> 1001 -> (2^3 * 1) + (2^2 * 0) + (2^1 * 0) + (2^0 * 1) -> 8 + 0 + 0 + 1 = 9

func Learn(run bool) {
	if !run {
		return
	}

	fmt.Println("Learning bit manipulation in Go")

	a := 9  // Binary: 1001
	b := 12 // Binary: 1100

	fmt.Println(a & b) // Bitwise AND: 1001 & 1100 = 1000 (8)
	fmt.Println(a | b) // Bitwise OR: 1001 | 1100 = 1101 (13)
	fmt.Println(a ^ b) // Bitwise XOR: 1001 ^ 1100 = 0101 (5)

	num := 4 // Binary: 0100
	// Shift left by 1 (4 * (2^1) = 8)
	leftShifted := num << 1 // Becomes 1000 (binary)
	fmt.Println(leftShifted)

	// Shift right by 1 (4 / (2^1) = 2)
	rightShifted := num >> 1 // Becomes 0010 (binary)
	fmt.Println(rightShifted)

	a1 := 2 ^ 3 // Bitwise XOR operator in Go not power of operator
	fmt.Println(a1)
	fmt.Println("Left Shifted by 3:", num<<3, num*int(math.Pow(2.0, 3.0)))
	num = num << 3
	fmt.Println("Right Shifted by 3:", num>>3, num/int(math.Pow(2.0, 3.0)))

	num1 := uint(0)
	fmt.Println(^num1>>1, int(math.Pow(2.0, 64-1)-1.0), uint((math.Pow(2.0, 64-1) - 1.0)), uint(b-a))

	bytess := []byte{
		'A': 65,
		'B': 66,
		'C': 67,
		'D': 68,
	}

	fmt.Println(bytess)
}
