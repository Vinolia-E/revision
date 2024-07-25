/*Instructions

Write a function that takes two integers and returns a string showing the range of numbers from the first to the second.

    The numbers must be separated by a comma and a space.
    If any of the arguments is bigger than 99 or less than 0, the function returns Invalid followed by a newline \n.
    Prepend a 0 to any number that is less than 10.
    Add a new line \n at the end of the string.

Expected function

func FromTo(from int, to int) string {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"psicine"
)

func main() {
	fmt.Print(piscine.FromTo(1, 10))
	fmt.Print(piscine.FromTo(10, 1))
	fmt.Print(piscine.FromTo(10, 10))
	fmt.Print(piscine.FromTo(100, 10))
}

And its output:

$ go run . | cat -e
01, 02, 03, 04, 05, 06, 07, 08, 09, 10$
10, 09, 08, 07, 06, 05, 04, 03, 02, 01$
10$
Invalid$*/

package main

import "fmt"

func Itoa(n int) string {
	//mod := n%10
	result := ""
	for n > 0 {
		mod := n%10
		result = string(rune(mod) + '0') + result
		n = n/10 
	}
	return result
}

func FromTo(from int, to int) string {
	word := ""
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}
	if from == to {
		word = Itoa(from)
	}
	
	//slice := []string{}


	if from > to {
		for i := from; i >= to; i-- {
			//fmt.Print(i)
			word = word + Itoa(i) + ", "
			
		}
		//return word + "\n"
	} else if from < to {
		for i := from; i <= to; i++ {
			word = word + Itoa(i) + ", "
		}
	}
	return word + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}

