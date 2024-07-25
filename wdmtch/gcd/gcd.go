// Write a function that takes two uint representing two strictly positive integers and returns their greatest common divisor. If any of the input numbers is 0, the function should return 0.

//     In mathematics, the greatest common divisor (GCD) of two or more integers, which are not all zero, is the largest positive integer that divides each of the integers.

// Expected function

// func Gcd(a, b uint) uint {

// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Gcd(42, 10))
// 	fmt.Println(piscine.Gcd(42, 12))
// 	fmt.Println(piscine.Gcd(14, 77))
// 	fmt.Println(piscine.Gcd(17, 3))
// }
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}

func Gcd(a, b uint) uint {
	var num uint//num := 0
	if a < b {
		num = a
	}else {
		num = b
	}
	for i := num; i >0; i-- {
		if a %i == 0 && b% i == 0 {
			return i
		} 
	}
return 0

}
