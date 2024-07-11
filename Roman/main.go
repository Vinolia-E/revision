package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
	//"strconv"
)

func main() {
	args := os.Args[1]
	if len(os.Args) != 2 {
		os.Exit(0)
	}

	num := Atoi(args) //, _ := strconv.Atoi(args)

	if num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit.")
		os.Exit(0)
	}

	result, calc := Rm(num)

	for _, ch := range result {
		z01.PrintRune(ch)
	}
	
	z01.PrintRune('\n')
	
	for _, chr := range calc {
		z01.PrintRune(chr)
	}
	z01.PrintRune('\n')
	// fmt.Println(result)
	// fmt.Println(calc)
}

func Rm(n int) (string, string) {
	var calc, result string

	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	cal := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "XL", "X", "(X-I)", "V", "(V-I)", "I"}

	for i, ch := range values {
		for n >= ch {
			n = n - ch
			result = result + roman[i]

			if len(calc) != 0 {
				calc += "+"
			}
			calc += cal[i]
		}
	}
	return calc, result
}

func Atoi(s string) int {
	number := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			fmt.Println("ERROR: cannot convert to roman digit.")
			os.Exit(0)
		} else {
			number = number*10 + int(ch-'0')
		}
	}
	return number
}
