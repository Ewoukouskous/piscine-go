package piscine

func IsAlpha(s string) bool {
	test := false
	for i := 0; i < len(s); i++ {
		if rune(s[i]) > 96 && rune(s[i]) < 123 || rune(s[i]) > 47 && rune(s[i]) < 58 || rune(s[i]) > 64 && rune(s[i]) < 91 {
			test = true
		} else {
			return false
		}
	}
	return test
}
