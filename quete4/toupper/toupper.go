package piscine

func ToUpper(s string) string {
	resultat := ""
	for i := 0; i < len(s); i++ {
		if rune(s[i]) > 96 && rune(s[i]) < 123 {
			newRune := rune(s[i]) - 32
			resultat += string(rune(newRune))
		} else {
			resultat += string(rune(s[i]))
		}
	}
	return resultat
}
