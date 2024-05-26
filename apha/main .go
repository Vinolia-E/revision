package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]
	for _, word := range args {
		if word%2 == 0 {
			z01.PrintRune(word - 32)
		} else {
			z01.PrintRune(word)
		}
	}
	z01.PrintRune('\n')
}
