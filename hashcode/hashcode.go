/* Instructions

Write a function called HashCode() that takes a string as an argument and returns a new hashed string.

    The hash equation is computed as follows:

(ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.

    If the resulting character is unprintable add 33 to it.

Expected function

func HashCode(dec string) string {
}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

And its output:

$ go run .
B
CD
EDF
Spwwz+bz}wo*/

package main

import "fmt"

//ASCII of current character + size of the string) % 127
func HashCode(dec string) string {
	result := ""
	new := 0

	for _, ch := range dec {
		new = (int(ch) + len(dec))%127
		// if new > 127 || new < 32 {
		// 	new += 33
		// }
		result += string(new)
	}
	return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
