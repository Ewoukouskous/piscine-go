package hackaton

func Rot14(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			if s[i]+14 <= 'Z' {
				res += string(s[i] + 14)
			} else {
				res += string('A' + (s[i] + 13) - 'Z')
			}
		} else if s[i] >= 'a' && s[i] <= 'z' {
			if s[i]+14 <= 'z' {
				res += string(s[i] + 14)
			} else {
				res += string('a' + (s[i] + 13) - 'z')
			}
		} else {
			res += string(s[i])
		}
	}
	return res
}
