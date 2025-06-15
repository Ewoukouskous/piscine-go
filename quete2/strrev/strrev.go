package piscine

func StrRev(s string) string {
	stri := ""
	for c := len(s) - 1; c >= 0; c-- {
		stri += string(s[c])
	}
	return stri
}
