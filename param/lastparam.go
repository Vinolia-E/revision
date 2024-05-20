package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	count := 0
	// arg  := args[len(args)-1] // last rune
	arg := args[0] // first rune

	for _, x := range arg {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')

	
		for _, chr := range args {
			for range chr {
				count++
			}
		}

	str := strconv.Itoa(count)
	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
