package piscine

func Capitalize(s string) string {
	first := true
	str := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' && first == false {
			str += string(rune(char) + 32)
		} else if char >= 'A' && char <= 'Z' {
			str += string(char)
			first = false
		} else if char >= 'a' && char <= 'z' && first == true {
			str += string(rune(char) - 32)
			first = false
		} else if char >= 'a' && char <= 'z' {
			str += string(char)
		} else if char >= '0' && char <= '9' {
			str += string(char)
			first = false
		} else {
			str += string(char)
			first = true
		}
	}
	return str
}
