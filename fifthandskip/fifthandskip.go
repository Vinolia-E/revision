/* Instructions

Write a function FifthAndSkip() that takes a string and returns another string. The function separates every five characters of the string with a space and removes the sixth one.

    If there are spaces in the middle of a word, ignore them and get the first character after the spaces until you reach a length of 5.
    If the string is less than 5 characters return Invalid Input followed by a newline \n.
    If the string is empty return a newline \n.

Expected function

func FifthAndSkip(str string) string {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(piscine.FifthAndSkip("This is a short sentence"))
	fmt.Print(piscine.FifthAndSkip("1234"))
}

And its output:

$ go run . | cat -e
abcde ghijk mnopq stuwx z$
Thisi ashor sente ce$
Invalid Input$
*/

package main

import (
	"fmt"
	"os"
)

func FifthAndSkip(str string) string {
	if len(str) < 5 {
		fmt.Println("Invalid Input")
		os.Exit(0)
	}

	word := ""

	for _, ch := range str {
		if ch == ' ' {
			continue
		}
		word += string(ch)
	}
	// return word + "\n"
	result := ""
	for i := 0; i < len(word); i += 5 {
		if i+5 > len(word) {
			result += word[i:]
		} else {
			result = result + word[i:i+5] //+ " "
			result += " "
		}
		i++
	}
	return result + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuvwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentencethe"))
	fmt.Print(FifthAndSkip("1234"))
}
