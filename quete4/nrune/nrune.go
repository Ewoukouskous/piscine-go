package piscine

func NRune(s string, n int) rune {
	if n >= 1 && n <= len(s) {
		return ([]rune(s))[n-1]
	}
	return 0
}
