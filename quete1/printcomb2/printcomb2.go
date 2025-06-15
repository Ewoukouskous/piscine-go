package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for w := 48; w < 58; w++ {
		for x := 48; x < 58; x++ {
			for y := 48; y < 58; y++ {
				for z := 48; z < 58; z++ {
					if w == 57 && x == 56 && y == 57 && z == 57 {
						z01.PrintRune(rune(w))
						z01.PrintRune(rune(x))
						z01.PrintRune(rune(32))
						z01.PrintRune(rune(y))
						z01.PrintRune(rune(z))
						z01.PrintRune(rune('\n'))
					} else {
						if w < y || w == y && x < z {
							z01.PrintRune(rune(w))
							z01.PrintRune(rune(x))
							z01.PrintRune(rune(32))
							z01.PrintRune(rune(y))
							z01.PrintRune(rune(z))
							z01.PrintRune(rune(44))
							z01.PrintRune(rune(32))
						}
					}
				}
			}
		}
	}
}
