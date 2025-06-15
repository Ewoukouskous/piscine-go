package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	var nbr []int
	if n == 0 {
		z01.PrintRune('0')
	}
	for n != 0 {
		if n < 0 {
			nbr = append(nbr, ((n%10)*(-1))+48)
		} else {
			nbr = append(nbr, (n%10)+48)
		}
		n /= 10
	}
	for k := 1; k < len(nbr)+1; k++ {
		z01.PrintRune(rune(nbr[len(nbr)-k]))
	}
}
