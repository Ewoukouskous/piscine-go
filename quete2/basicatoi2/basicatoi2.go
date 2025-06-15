package piscine

func BasicAtoi2(s string) int {
	num := 0
	for c := 0; c < len(s); c++ {
		if rune(s[c]) != 48 && rune(s[c]) != 49 && rune(s[c]) != 50 && rune(s[c]) != 51 && rune(s[c]) != 52 && rune(s[c]) != 53 && rune(s[c]) != 54 && rune(s[c]) != 55 && rune(s[c]) != 55 && rune(s[c]) != 56 && rune(s[c]) != 57 {
			return 0
		} else {
			num *= 10
			num += int(rune(s[c])) - 48
		}
	}
	return num
}
