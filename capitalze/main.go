/*
Instructions

Write a function that capitalizes the first letter of each word and lowercases the rest.

A word is a sequence of alphanumeric characters.
Expected function

func Capitalize(s string) string {

}

Usage

Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
}

And its output :

$ go run .
Hello! How Are You? How+Are+Things+4you?
$
*/

package main

import "fmt"

func Capitalize(s string) string {
	var res rune
	result := ""
	caps := false
	for _, char := range s {
		if caps {
			if char >= 'a' && char <= 'z' {
				res = char - 'a' + 'A'
				caps = false
			} else if char >= 'A' && char <= 'Z'  {
				res =  char - 'A' + 'a'
				caps = true
			} else {
				res = char
			}
			result += string(res)
		}
	}
	return result
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
