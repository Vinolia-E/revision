/*Instructions

Write a program that takes a string and displays it with exactly three spaces between each word, with no spaces nor tabs at neither the beginning nor the end.

The string will be followed by a newline ('\n').

A word, in this exercise, is a sequence of visible characters.

If the number of arguments is not 1, or if there are no word, the program displays nothing.
Usage

$ go run . "you   see   it's   easy   to   display   the   same   thing" | cat -e
you   see   it's   easy   to   display   the   same   thing$
$ go run . "   only  it's harder   " | cat -e
only   it's   harder$
$ go run . " how funny it is" "did you  hear, Mathilde ?" | cat -e
$ go run .
$
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]
	splited := Split(str)
	
	result := ""
	for _, word := range splited {
		if word == " " {
			continue
		} else {
			result += word + "  "
		}
	}

	fmt.Println(result[:len(result)-2])
}

func Split(str string) []string {
	result := ""
	word := []string{}

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			result += string(str[i])
		} else if str[i] == ' ' && result != "" {
			word = append(word, result)
			result = ""
		}
	}
	if result != "" {
		word = append(word, result)
	}
	return word
}
