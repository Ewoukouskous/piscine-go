package main

import "github.com/01-edu/z01"

func setPoint(x, y *int) {
	*x = 42
	*y = 21
}

func main() {
	a := 0
	b := 0
	setPoint(&a, &b)

	z01.PrintRune(120)
	z01.PrintRune(32)
	z01.PrintRune(61)
	z01.PrintRune(32)
	z01.PrintRune(52)
	z01.PrintRune(50)
	z01.PrintRune(44)
	z01.PrintRune(32)
	z01.PrintRune(121)
	z01.PrintRune(32)
	z01.PrintRune(61)
	z01.PrintRune(32)
	z01.PrintRune(50)
	z01.PrintRune(49)
	z01.PrintRune('\n')
}
