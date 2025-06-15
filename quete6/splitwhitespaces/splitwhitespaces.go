package splitwhitespaces

func SplitWhiteSpaces(s string) []string {
	tab := make([]string, 0)
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if i > start {
				tab = append(tab, s[start:i])
			}
			start = i + 1
		}
	}
	if start < len(s) {
		tab = append(tab, s[start:])
	}
	return tab
}
