package main

import "github.com/01-edu/z01"

func main() {
	for i := 'z'; i >= 'a'; i-- {
		if i %2 == 1 {
			j := i-32
			z01.PrintRune(j)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}