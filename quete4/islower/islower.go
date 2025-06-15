package piscine

func IsLower(s string) bool {
	test := false
	for i := 0; i < len(s); i++ {
		if rune(s[i]) > 96 && rune(s[i]) < 123 {
			test = true
		} else {
			return false
		}
	}
	return test
}
