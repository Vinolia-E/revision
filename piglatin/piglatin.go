package main

import (
	"os"
	"fmt"
) 

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]

	piglatin := Piglatin(word)
	fmt.Println(piglatin)
}

func Piglatin(str string) string {
	vowels := "aeiouAEIOU"
	latin := "" 

	for i, ch := range str {
		for _, vowel := range vowels {
			if i == 0 && ch == vowel {
				latin = str + "ay"
				return latin
			} else if ch == vowel {
				latin = str[i:] + "ay" + str[:i]
				return latin
			} 
		}
	}
	return ""
}