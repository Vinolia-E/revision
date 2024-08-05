/*wordflip
Instructions

Write a function WordFlip() that takes a string as input and returns it in reverse order.

    The output should be followed by a newline \n.
    If the string is empty, return Invalid Output.
    Ignore multiple spaces between words and trim any leading or trailing spaces in the string.

Expected function

func WordFlip(str string) string {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.WordFlip("First second last"))
	fmt.Print(piscine.WordFlip(""))
	fmt.Print(piscine.WordFlip("     "))
	fmt.Print(piscine.WordFlip(" hello  all  of  you! "))
}

And its output:

$ go run . | cat -e
last second First$
Invalid Output$
$
you! of all hello$*/

package main

import "fmt"

func WordFlip(str string) string {
	slice := []string{}
	var word, result string

	if len(str) < 1 {
		return "Invalid Input,\n"
	}
	for i := len(str)-1; i >= 0; i-- {
		if str[i] != ' ' {
			word =string(str[i]) + word
		} else if str[i] == ' ' && word != "" {
			slice = append(slice, word)
			word = ""
		}
		if i == 0 {
			slice = append(slice, word)
		}
	}

	for _, words := range slice {
		if result != "" {
			result += " "
		}
		result += words
	}
	if len(result) > 0 {
		
	}
	return result + "\n"
}

func main() {
	fmt.Print(
		WordFlip("First second last"))
	fmt.Print(
		WordFlip(""))
	fmt.Print(
		WordFlip("     "))
	fmt.Print(
		WordFlip(" hello  all  of  you! "))
}
