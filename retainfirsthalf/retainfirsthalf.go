/*nstructions

Write a function called RetainFirstHalf() that takes a string as an argument and returns the first half of this string.

    If the length of the string is odd, round it down.
    If the string is empty, return an empty string.
    If the string length is equal to one, return the string.

Expected function

func RetainFirstHalf(str string) string {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(piscine.RetainFirstHalf("A"))
	fmt.Println(piscine.RetainFirstHalf(""))
	fmt.Println(piscine.RetainFirstHalf("Hello World"))
}

And its output:

$ go run . | cat -e
This is the 1st half$
A$
$
Hello$
*/

package main

import "fmt"

func RetainFirstHalf(str string) string {
	result := ""
	length := len(str)
	if length == 0 {
		return ""
	} else if length == 1 {
		return str
	} else if length % 2 == 1 {
		length -= 1
	}
	half := length/2

	for i := 0; i < half; i++ {
		result += string(str[i])
	}
	return result
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
