/*Write a function IsCapitalized() that takes a string as an argument and returns true if each word in the string begins with either an uppercase letter or a non-alphabetic character.

    If any of the words begin with a lowercase letter return false.
    If the string is empty return false.

Expected function

func IsCapitalized(s string) bool {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
} */

package main

import "fmt"

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	sliced := SplitString(s)

	for _, word := range sliced {
		for i, chr := range word {
			if i == 0 && chr >= 'a' {
				return false
			}
		}
	}
	return true
}

func SplitString(s string) []string{
	str := ""
	strSlice := []string{}
	for i, ch := range s {
		if ch == ' ' {
			strSlice = append(strSlice, str)
			str = ""
			continue
		}

		if ch == ' ' {
			continue
		}
		 str += string(ch)
		if i == len(s)-1 {
			strSlice = append(strSlice, str)
		}

	}
	//fmt.Println(len(strSlice))
	return strSlice
}


func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
