package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for _, arg := range args[1:] {
		printString(arg)
		printString("\n")
	}
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
