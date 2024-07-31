/* Instructions

Write a function that takes (arr [10]byte), and displays the memory as in the example.

After displaying the memory the function must display all the ASCII graphic characters. The non printable characters must be replaced by a dot.

The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.
Expected function

func PrintMemory(arr [10]byte) {

}

Usage

Here is a possible program to test your function :

package main

import "piscine"

func main() {
	piscine.PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

And its output :

$ go run . | cat -e
68 65 6c 6c$
6f 10 15 2a$
00 00$
hello..*..$
$
*/

package main

import "fmt"

func PrintMemory(arr [10]byte) {
	result := []string{}
	for _, ch := range arr {
		if rune(ch) < 32 || rune(ch) >127 {
			result = append(result, "00")
		} else {
			result = append(result, Hex(int(ch)))
		}
	}
// fmt.Print(result)
// fmt.Print("\n")
for i := 0; i < len(result); i++ {
	fmt.Print(result[i])
	fmt.Print(" ")
	if i == 3 || i == 7 || i == 9 {
		fmt.Print("\n")
	}
}

str := ""
for _, ch := range arr {
	if ch < 32 || ch > 127 {
		str += "."
		} else {
			str += string(ch)
		}
	}
	fmt.Print(str,"\n")
}

func Hex(n int) string {
	hex := "0123456789abcdef"
	result := ""

	for n > 0 {
		mod := n%16
		result = string(hex[mod]) + result
		n /= 16
	}
	return result
}

func main() {
	
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}