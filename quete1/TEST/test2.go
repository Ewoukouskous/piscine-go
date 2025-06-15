package main

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		code := int(s[i])
		z01.PrintRune(rune(code))
	}
}

func main() {
	PrintStr("Hello World!")
	z01.PrintRune(rune('\n'))
}
