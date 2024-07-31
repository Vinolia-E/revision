/*Instructions

Write a function, Unmatch, that returns the element of the slice that does not have a correspondent pair.

    If all the number have a correspondent pair, it should return -1.

Expected function

func Unmatch(a []int) int {

}

Usage

Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := piscine.Unmatch(a)
	fmt.Println(unmatch)
}

And its output :

$ go run .
4
$
*/
