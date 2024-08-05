/*revconcatalternate
Instructions

Write a function RevConcatAlternate() that receives two slices of int as arguments and returns a new slice with alternated values of each slice in reverse order.

    The input slices can have different lengths.
    The new slice should start with the elements from the largest slice first and when they became equal size slices, it should add an element of the first given slice.
    If the slices are of equal length, the new slice should start with an element of the first slice.

    Note: you can check the examples bellow for more details.

Expected function

func RevConcatAlternate(slice1,slice2 []int) []int {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

And its output:

$ go run .
[3 6 2 5 1 4]
[9 8 7 3 6 2 5 1 4]
[8 9 3 2 5 1 4]
[3 2 1]*/

package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {
	new := []int{}
	length1 := len(slice1)
	length2 := len(slice2)
	i := len(slice1) - 1
	j := len(slice2) - 1

	// dif := len(slice1)- len(slice2)
	for i >= 0 && j >= 0 {
		if length1 > length2 {
			new = append(new, slice1[i])
			length1--
			i--

		} else if length2 > length1 {
			new = append(new, slice2[j])
			j--
			length2--
		} else {
			new = append(new, slice1[i])
			i--
			length1--
		}
	}
	for i >= 0 {
		new = append(new, slice1[i])
		i--
	}
	for j >= 0 {
		new = append(new, slice2[j])
		j--
	}

	return new
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

// func main() {
// 	// fmt.Println(
// 	// 	RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
// 	// fmt.Println(
// 	// 	RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(
// 		RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5,6,7}))
// 	// fmt.Println(
// 	// 	RevConcatAlternate([]int{1, 2, 3}, []int{}))
// }
