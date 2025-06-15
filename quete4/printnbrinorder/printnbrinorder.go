package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n <= 0 {
		z01.PrintRune(48)
	}

	//if n >= 4192470505494085389 {
	//	return
	//}

	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	for i := 0; i < len(digits); i++ {
		minIndex := i
		for j := i + 1; j < len(digits); j++ {
			if digits[j] < digits[minIndex] {
				minIndex = j
			}
		}
		digits[i], digits[minIndex] = digits[minIndex], digits[i]
	}

	for _, d := range digits {
		z01.PrintRune(rune('0' + d))
	}
}
