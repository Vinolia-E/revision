package main

import "fmt"

func main() {
	s1 := "hella"
	s2 := "a"
	s3 := "o"

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