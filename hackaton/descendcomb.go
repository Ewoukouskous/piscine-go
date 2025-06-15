package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 57; i > 47; i-- {
		for j := 57; j > 47; j-- {
			for k := 57; k > 47; k-- {
				for l := 57; l > 47; l-- {
					if i > k {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(32))
						z01.PrintRune(rune(k))
						z01.PrintRune(rune(l))
						if i == 48 && j == 49 && k == 48 && l == 48 {
						} else {
							z01.PrintRune(rune(44))
							z01.PrintRune(rune(32))
						}
					} else if i == k && j > l {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(32))
						z01.PrintRune(rune(k))
						z01.PrintRune(rune(l))
						if i == 48 && j == 49 && k == 48 && l == 48 {
						} else {
							z01.PrintRune(rune(44))
							z01.PrintRune(rune(32))
						}
					}
				}
			}
		}
	}
}
