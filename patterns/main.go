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

func Pattern11() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 5; j++ {
			condition := (i <= 5 && j <= i) || (i > 5 && j <= 5-(i-5))

			if condition {
				fmt.Printf("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern12() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 9; j++ {
			condition := j >= i && j <= 10-i

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

func Pattern13() {
	for i := 1; i <= 5; i++ {
		k := i
		for j := 1; j <= 9; j++ {
			condition := j >= 6-i && j <= 4+i

			if condition {
				fmt.Printf("%v ", k)

				if j < 5 {
					k++
				} else {
					k--
				}
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern14() {
	for i := 1; i <= 9; i++ {
		k := 9 - i
		for j := 1; j <= 9; j++ {
			condition := j <= 10-i

			if condition {
				fmt.Printf("%v ", k)
				k--
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern15() {
	for i := 1; i <= 9; i++ {
		k := 1
		for j := 1; j <= 5; j++ {
			condition := (i <= 5 && j >= 6-i) || (i > 5 && j >= i-4)

			if condition {
				fmt.Printf("%v ", k)
				k++
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern16() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			switch {
			case i == j:
				fmt.Print("\\")
			case 10-i == j:
				fmt.Print("/")
			default:
				fmt.Print("*")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern17() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			condition := (i <= 5 && ((j <= 6-i) || (j >= 4+i))) || (i > 5 && ((j <= i-4) || (j >= 9-(i-5))))

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

func Pattern18() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 9; j++ {
			condition := j >= i && j <= 10-i

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

func Pattern19() {
	for i := 1; i <= 5; i++ {
		c := uint8(65)
		for j := 1; j <= 9; j++ {
			condition := j >= 6-i && j <= 4+i

			if condition {
				fmt.Printf("%v ", string(c))
				if j < 5 {
					c++
				} else {
					c--
				}
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func Pattern20() {
	for i := 1; i <= 4; i++ {
		c := uint8(65)
		n := uint8(1)
		for j := 1; j <= 8; j++ {
			condition := j >= 5-i && j <= 4+i

			if condition {
				if j < 5 {
					fmt.Printf("%v ", string(c))
					c++
				} else {
					fmt.Printf("%v ", n)
					n++
				}

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
	Pattern11()
	Pattern12()
	Pattern13()
	Pattern14()
	Pattern15()
	Pattern16()
	Pattern17()
	Pattern18()
	Pattern19()
	Pattern20()
}
