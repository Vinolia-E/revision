/*rostring
Instructions

Write a program that takes a string and displays this string after rotating it one word to the left.

Thus, the first word becomes the last, and others stay in the same order.

A word is a sequence of alphanumerical characters.

Words will be separated by only one space in the output.

If the number of arguments is different from 1, the program displays a newline.
Usage

$ go run . "abc   " | cat -e
abc$
$ go run . "Let there     be light"
there be light Let
$ go run . "     AkjhZ zLKIJz , 23y"
zLKIJz , 23y AkjhZ
$ go run . | cat -e
$
$
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	str := ""

	args := os.Args[1]

	revsplited := []string{}
	splited := Split(args)
	first := splited[0]
	rest := splited[1:]

	revsplited = append(revsplited, rest...)
	revsplited = append(revsplited, first)
	//fmt.Println(revsplited)

	for _, word := range revsplited {
		if str != "" {
			str += " "
		}
		str += word
	}
	fmt.Println(str)
}

func Split(s string) []string {
	word := ""
	slice := []string{}

	for i, ch := range s {
		if ch != ' ' {
			word += string(ch)
		} else if ch == ' ' && word != "" {
			slice = append(slice, word)
			word = ""
		}
		if i == len(s)-1 {
			slice = append(slice, word)
		}
	}
	return slice
}