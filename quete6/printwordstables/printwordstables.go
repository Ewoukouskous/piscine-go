package printwordstables

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if j == len(a[i])-1 {
				z01.PrintRune(rune(a[i][j]))
				z01.PrintRune(rune('\n'))
			} else {
				z01.PrintRune(rune(a[i][j]))
			}
		}
	}
}
