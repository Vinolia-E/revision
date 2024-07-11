package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]

	piglatin := Piglatin(word)
	if piglatin == "" {
		return
	}
	fmt.Println(piglatin)
}

func Piglatin(str string) string {
	vowels := "aeiouAEIOU"

	for i, ch := range str {
		for _, vowel := range vowels {
			if ch == vowel { // ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
				if i == 0 {
					return str + "ay"
				} else {
					return string(str[i:]) + "ay" + string(str[:i])
				}
			}
		}
	}
	return ""
}
