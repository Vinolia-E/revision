package main

import "github.com/01-edu/z01"

func main() {
	word := "hello world"
	for _, ch := range word {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}