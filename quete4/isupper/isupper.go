package piscine

func IsUpper(s string) bool {
	test := false
	for i := 0; i < len(s); i++ {
		if rune(s[i]) > 64 && rune(s[i]) < 91 {
			test = true
		} else {
			return false
		}
	}
	return test
}
