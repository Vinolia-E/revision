package main

import (
	"fmt"
	"os"
)

func main() {
	// s1 := "hella"
	// s2 := "a"
	// s3 := "o"
	s1 := os.Args[1]
	s2 := os.Args[2]
	s3 := os.Args[3]

	str := ""

	for _, word := range s1 {
		if string(word) == s2 {
			str += s3
		} else {
			str += string(word)
		}
	}
	fmt.Println(str)
}