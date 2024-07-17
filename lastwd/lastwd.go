// Write a program that takes a string and displays its last word, followed by a newline ('\n').

//     A word is a section of string delimited by spaces or by the start/end of the string.

//     The output will be followed by a newline ('\n').

//     If the number of arguments is different from 1, or if there are no word, the program displays nothing.

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	if len(args) != 2{
		return
	}

	str := args[1]
	check := false
	
	for _, ch := range str{
		if ch != ' '{
			check = true
			break
		}
	}
	if check {
		result := ""
		for i := len(str)-1; i >= 0; i--{
			if rune(str[i]) != ' ' {
				result = string(str[i]) + result
			 }else if result != ""{
				break
			 }
		}
		for _, c := range result{
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}

}