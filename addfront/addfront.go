/* Instructions

Write a function that takes a string and a slice of strings, this function will return a new slice of string with the given string prepended

    If the given string is empty you only need to return the given slice

Expected function

func AddFront(s string, slice []string) []string {
    // your code here
}

Usage

Here is a possible program to test your function:

package main

import "fmt"

func main() {
    fmt.Println(AddFront("Hello", []string{"world"}))
    fmt.Println(AddFront("Hello", []string{"world", "!"}))
    fmt.Println(AddFront("Hello", []string{}))
}

and the output should be:

$ go run .
[Hello world]
[Hello world !]
[Hello]
*/

package main

import "fmt"

func AddFront(s string, slice []string) []string {
	var result []string 
	result = append(result, s)
	for _, ch := range slice {
    result = append(result, ch)
	}
	return result
}

func main() {
    fmt.Println(AddFront("Hello", []string{"world"}))
    fmt.Println(AddFront("Hello", []string{"world", "!"}))
    fmt.Println(AddFront("Hello", []string{}))
}
