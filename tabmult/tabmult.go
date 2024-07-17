// Write a program that displays a number's multiplication table.

//     The parameter will always be a strictly positive number that fits in an int. Said number multiplied by 9 will also fit in an int.

// Usage

// $ go run . 9
// 1 x 9 = 9
// 2 x 9 = 18
// 3 x 9 = 27
// 4 x 9 = 36
// 5 x 9 = 45
// 6 x 9 = 54
// 7 x 9 = 63
// 8 x 9 = 72
// 9 x 9 = 81
// $ go run . 19
// 1 x 19 = 19
// 2 x 19 = 38
// 3 x 19 = 57
// 4 x 19 = 76
// 5 x 19 = 95
// 6 x 19 = 114
// 7 x 19 = 133
// 8 x 19 = 152
// 9 x 19 = 171
// $ go run .

// $

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main(){
	args := os.Args
	 if len(args) != 2 {
		return
	 }
	 str :=args[1]
	 num := 0

	 for _, ch := range str{
		if ch >= '0' && ch <= '9'{
			num = num * 10 + int(ch - '0')
		}else {
			return
		}
	 }
	 result := 0
	 for i := 1; i <= 9; i++ {
		result = i * num
		z01.PrintRune(rune(i + '0'))
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		Itoa(num)
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		Itoa(result)
		z01.PrintRune('\n')
	 }
}

func Itoa(nb int){
	result := ""
	for nb != 0{
		mod := nb % 10
		start := '0'
		for i := 0; i < mod; i++{
			start++
		}
		result = string(start) + result
		nb = nb / 10
	}
	for _, ch := range result {
		z01.PrintRune(ch)
	}
}