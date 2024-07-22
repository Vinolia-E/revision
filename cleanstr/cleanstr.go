/*Instructions

Write a program that takes a string, and displays this string with exactly:

    one space between words.
    without spaces nor tabs at the beginning nor at the end.
    with the result followed by a newline ("\n").

A "word" is defined as a part of a string delimited either by spaces/tabs, or by the start/end of the string.

If the number of arguments is not 1, or if there are no words to display, the program displays a newline("\n").
Usage :

$ go run . "you see it's easy to display the same thing" | cat -e
you see it's easy to display the same thing$
$ go run . " only    it's  harder   "
only it's harder$
$ go run . " how funny" "Did you   hear Mathilde ?"

$ go run . ""

$ */

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("\n")
	}

	args := os.Args[1]
	word := ""
	sliced := []string{}

	for i, ch := range args {
		if ch == ' ' {
			sliced = append(sliced, word)
			word = ""
			continue
		}
		word += string(ch)
		if i == len(args)-1 {
			sliced = append(sliced, word)
		}

	}
	fmt.Println(len(sliced))
	fmt.Println(sliced)
	str := ""
	for _, words := range sliced {
		if words != "" {
			// str += " "
			str += words + " "
			// continue
		}
	}
	if str[len(str)-1] == ' ' {
		fmt.Println(str[:(len(str) - 1)])
	} else {
		fmt.Println(str)
	}
}
