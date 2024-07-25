/*Instructions

Write a function that returns the median of five int arguments.
Expected function

func Abort(a, b, c, d, e int) int {

}

Usage

Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	middle := piscine.Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}

And its output :

$ go run .
5
$
*/

package main 

import "fmt"

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}

func Abort(a, b, c, d, e int) int {

	var slice [] int
	slice = append(slice, a, b, c, d, e)

	for i := 0; i < len(slice); i++ {
		for j := i+1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice[2]
}
