package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	args := os.Args[1]
	res := Split(args)
	result := ""
	
	for _, ch := range res {
		if ch == "" {
			result += ch
		} else {
			result += ch + " "
		}
	}
	fmt.Println(result)
}

func Split(str string) []string {
	result := ""
	word := []string{}

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			result += string(str[i])
		} else if str[i] == ' ' && result != "" {
			word = append(word, result)
			result = ""
		}
	}
	if result != "" {
		word = append(word, result)
	}
	return word
}
