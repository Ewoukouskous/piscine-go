package piscine

func IsPrintable(s string) bool {
	test := false
	for i := 0; i < len(s); i++ {
		if rune(s[i]) > 31 && rune(s[i]) < 127 {
			test = true
		} else {
			return false
		}
	}
	return test
}
