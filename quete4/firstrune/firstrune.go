package piscine

func FirstRune(str string) rune {
	for _, r := range str {
		return rune(r)
	}
	return 1
}
