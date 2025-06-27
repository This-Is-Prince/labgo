package main

import "fmt"

func Pattern1() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j <= i {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern2() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			condition := j >= 6-i
			if condition {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern3() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			condition := j >= i
			if condition {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern4() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			condition := j <= 6-i
			if condition {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern5() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 9; j++ {
			condition := j >= 6-i && j <= 4+i
			if condition {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern6() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 9; j++ {
			condition := j >= 6-i && j <= 4+i && ((i%2 == 0 && j%2 == 0) || (i%2 != 0 && j%2 != 0))
			if condition {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern7() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 9; j++ {
			condition := j <= 6-i || j >= 4+i
			if condition {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern8() {
	for i := 1; i <= 5; i++ {
		a := 1
		for j := 1; j <= 9; j++ {
			condition := j >= 6-i && j <= 4+i
			if condition {
				fmt.Print(a, " ")

				if j < 5 {
					a++
				} else {
					a--
				}
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern9() {
	for i := 1; i <= 5; i++ {
		a := uint8(65)
		for j := 1; j <= 9; j++ {
			condition := j <= 6-i || j >= 4+i
			if condition {
				fmt.Print(string(a), " ")

				if j < 6-i {
					a++
				} else if j >= 4+i {
					a--
				}
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern10() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			condition := (6-i > 0 && j >= 6-i && j <= 4+i) || (6-i <= 0 && j >= i-4 && j <= 9-(i-5))
			if condition {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func main() {
	fmt.Println("Patterns in Go!")

	Pattern1()
	Pattern2()
	Pattern3()
	Pattern4()
	Pattern5()
	Pattern6()
	Pattern7()
	Pattern9()
	Pattern10()
}
