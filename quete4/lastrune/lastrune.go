package main

import "github.com/01-edu/z01"

func LastRune(str string) rune {
	return rune(str[len(str)-1])
}

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}
