/* Instructions

Write a function that takes two strings's and returns the number of characters that are not included in both, without repeating characters.

    If there is no unique characters return 0.
    If both strings are empty return -1.

Expected function

func WeAreUnique(str1 , str2 string) int {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

And its output:

$ go run .
2
-1
6*/

package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
	fmt.Println(WeAreUnique("", "banana"))
}

func WeAreUnique(str1 , str2 string) int {
	count := 0
	result := ""
	if len(str1)== 0 && len(str2) == 0 {
		return -1
	}
	for _, ch := range str1 {
		if !strings.Contains(str2, string(ch)){
			if strings.Contains(result, string(ch)) {
				continue
			}
			result += string(ch)
			count++
		}
	}

	for _, c := range str2 {
		if !strings.Contains(str1, string(c)){
			if strings.Contains(result, string(c)) {
				continue
			}
			result += string(c)
			count++
		}
	}
	 fmt.Println(result)
	return count
}