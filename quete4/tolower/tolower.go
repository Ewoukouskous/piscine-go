package piscine

func ToLower(s string) string {
	resultat := ""
	for i := 0; i < len(s); i++ {
		if rune(s[i]) > 64 && rune(s[i]) < 91 {
			newRune := rune(s[i]) + 32
			resultat += string(rune(newRune))
		} else {
			resultat += string(rune(s[i]))
		}
	}
	return resultat
}
