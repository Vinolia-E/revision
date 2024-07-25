/*Instructions

Write a function AdvancedSortWordArr that sorts a slice of string, based on the function f passed in parameter.
Expected function

func AdvancedSortWordArr(a []string, f func(a, b string) int) {

}

Usage

Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	piscine.AdvancedSortWordArr(result, piscine.Compare)

	fmt.Println(result)
}

And its output :

$ go run .
[1 2 3 A B C a b c]
$
*/

package main

import "fmt"

func AdvancedSortWordArr(a []string, f func(a, b string) int) {

	//var result []string
	for _, ch := range a {
		for f(ch) {

		}
	}

}

func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	AdvancedSortWordArr(result, Compare)

	fmt.Println(result)
}

func Compare(a, b string) int {
	if a == b {
		return 0 
	} else if a > b {
		return 1
	} else{
		return -1
	}
}