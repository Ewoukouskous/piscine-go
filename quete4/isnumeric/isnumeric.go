package piscine

func IsNumeric(s string) bool {
	test := false
	for i := 0; i < len(s); i++ {
		if rune(s[i]) > 47 && rune(s[i]) < 58 {
			test = true
		} else {
			return false
		}
	}
	return test
}
