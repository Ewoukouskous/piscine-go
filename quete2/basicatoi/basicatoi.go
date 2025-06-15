package piscine

func BasicAtoi(s string) int {
	num := 0
	for c := 0; c < len(s); c++ {
		num *= 10
		num += int(rune(s[c])) - 48
	}
	return num
}
