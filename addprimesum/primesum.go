/*addprimesum
Instructions

Write a program that takes a positive integer as argument and displays the sum of all prime numbers inferior or equal to it followed by a newline ('\n').

    If the number of arguments is different from 1, or if the argument is not a positive number, the program displays 0 followed by a newline.

Usage

$ go run . 5
10
$ go run . 7
17
$ go run . -2
0
$ go run . 0
0
$ go run .
0
$ go run . 5 7
0
$
*/

package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {

	result := 0
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	args := os.Args[1]

	number := ToInt(args)

	if number < 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	} else {
		if number > 0 {
			for i := 2; i <= number; i++ {
				if Isprime(i) {
					result += i
				}
			}
		}
	}
	fmt.Println(result)
}

func ToString(n int) string {
	mod := n%10
	str := ""

	for n > 0 { 
		str = string(rune(mod+'0')) + str
		n /= 10
	}
	return str

}

func ToInt(s string) int {
	sign := 1
	num := 0

	for i, ch := range s {
		if i == 0 && ch == '-' {
			sign = -1
			continue
		} else if i == 0 && ch == '+' {
			sign = 1
			continue
		}
		if ch >= '0' && ch <= '9' {
			num = num *10 + int(ch-'0')
		}
	}
	return num * sign
}
func Isprime(n int) bool {
	for i := 2; i < n; i++ {
		if n %i == 0 {
			return false
		} 
	}
	return true
}