package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for i := len(args) - 1; i > -1; i-- {
		printString(args[i])
		printString("\n")
	}
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
