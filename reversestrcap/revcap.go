/*reversestrcap
Instructions

Write a program that takes one or more arguments and that, for each argument, puts the last letter of each word in uppercase and the rest in lowercase. It displays the result followed by a newline ('\n').

If there are no argument, the program displays nothing.
Usage

$ go run . "First SMALL TesT" | cat -e
firsT smalL tesT$
$ go run . "SEconD Test IS a LItTLE EasIEr" "bEwaRe IT'S NoT HARd WhEN " " Go a dernier 0123456789 for the road e" | cat -e
seconD tesT iS A littlE easieR$
bewarE it'S noT harD wheN $
 gO A dernieR 0123456789 foR thE roaD E$
$ go run .
$
*/

package main

import (
	"fmt"
	"os"

)

func main() {
	if len(os.Args) < 2 {
		return
	}

	args := os.Args[1:]
	result := ""

	for _, arg := range args {
		splited := Split(arg)
		for _, word := range splited {
			for i, ch := range word {
				if Islowwer(ch) && i == len(word)-1 {
					ch -= 32
					result += string(ch)
				} else if IsUpper(ch) && i != len(word)-1 {
					ch = Tolowwer(ch)
					result += string(ch)
				} else {
					result += string(ch)
				}
			}
		}
		fmt.Println(result)
		result = ""
//fmt.Println(Split(arg))
//fmt.Println(string(Tolowwer('2')))
	}
	// fmt.Println(result)
}

func Split(s string) []string {
	word := ""
	slice := []string{}

	for i, ch := range s {
		if ch != ' ' {
			word += string(ch)
		} else if ch == ' ' && word != "" {
			slice = append(slice, " ", word)
			word = ""
		}
		if i == len(s)-1 {
			slice = append(slice, " ", word)
		}
	}
	return slice
}

func IsUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func Tolowwer(r rune) rune{
	if r >= 'A' && r <= 'Z' {
		 r += 32
	}
	return r
}

func Islowwer(r rune) bool {
	return r >='a'&& r <= 'z'
}