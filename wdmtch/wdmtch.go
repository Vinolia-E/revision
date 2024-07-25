package main

import (
	//"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	s1 := args[0]
	s2 := args[1]
	index := 0
	if len(args) != 2 {
		return
	}

	result := ""

	for _, ch := range s2 {
		for i := index; i < len(s1); i++ {
			if ch == rune(s1[index]) {
				result += string(ch)
				index++
				break
			}
		}
	}
	 if len(result) == len(s1) {
		for _, ch := range result {
			z01.PrintRune(ch)
		}
	 	//fmt.Println(result)
	 } else {
		return
	 }
	z01.PrintRune('\n')
}