package main

import (
	"fmt"
)

func StrLen(s string) int {
	return len(s)
// 	count := 0
// 	for range s {
// 		count++
// 	}
// 	return count
 }

func main() {
	l := StrLen("Hello there beautiful World!")
	fmt.Println(l)
}
