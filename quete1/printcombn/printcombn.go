package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	digits := "0123456789"
	printCombinations(digits, "", n, true)
	z01.PrintRune(rune(10))
}

func printCombinations(digits, prefix string, n int, first bool) {
	if n == 0 {
		if !first {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		for _, r := range prefix {
			z01.PrintRune(rune(r))
		}
		return
	}

	for i := 0; i < len(digits); i++ {
		newPrefix := prefix + string(digits[i])
		remainingDigits := digits[i+1:]

		printCombinations(remainingDigits, newPrefix, n-1, first && i == 0)
	}
}
