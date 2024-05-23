package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	s1 := args[0]
	s2 := []rune(args[1])

	count := 0
	str := ""

	for _, cha := range s1 {
		for i := count; i < len(s2); i++ {
			text := s2[i]
			if text == cha {
				str += string(text)
				count = i+1
				break
			}
		}
	}
	if str == s1 {
		for _, ch := range str {
			z01.PrintRune(ch)
			z01.PrintRune('\n')
		}
	}
	//z01.PrintRune('\n')
}