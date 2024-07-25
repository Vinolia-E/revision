/* Instructions

Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

    Factors must be displayed in ascending order and separated by *.

    If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.

Usage

$ go run . 225225
3*3*5*5*7*11*13
$ go run . 8333325
3*3*5*5*7*11*13*37
$ go run . 9539
9539
$ go run . 804577
804577
$ go run . 42
2*3*7
$ go run . a
$ go run . 0
$ go run . 1
$
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]
	number := Atoi(args)
	if number < 2 {
		return
	}

	result := ""

	for i := 2; i <= number; {
		if Isprime(i) && number%i == 0 {
			if result != "" {
				result += "*"
			}
			result += Itoa(i) 
			number = number/i
			i =2
		}
		i++
	}
	fmt.Println(result)
}

func Isprime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Atoi(s string) int{
	result := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			result = result * 10 + int(ch - '0')
		}
	}
	return result
}

func Itoa(n int) string {
	result := ""
	for n > 0 {
		mod := n % 10
		result = string(rune(mod)+ '0')
		n = n/10
	}
	return result
}