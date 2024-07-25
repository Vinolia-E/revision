package main
// Write a program that transforms a string passed as argument in its Pig Latin version.

// The rules used by Pig Latin are as follows:

//     If a word begins with a vowel, just add "ay" to the end.
//     If it begins with a consonant, then we take all consonants before the first vowel and we put them on the end of the word and add "ay" at the end.
//     Only the latin vowels will be considered as vowel(aeiou).

// If the word has no vowels, the program should print "No vowels".

// If the number of arguments is different from one, the program prints nothing.
// Usage

// $ go run .
// $ go run . pig | cat -e
// igpay$
// $ go run . Is | cat -e
// Isay$
// $ go run . crunch | cat -e
// unchcray$
// $ go run . crnch | cat -e
// No vowels$
// $ go run . something else | cat -e
// $


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
	return "No vowel"
}