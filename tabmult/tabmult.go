// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main(){
// 	args := os.Args
// 	if len(args) != 2 {
// 		return
// 	}
// 	said := args[1]
// 	num := 0
// 	for _, ch := range said {
// 		if ch >= 0 || ch <= 9 {
// 			num = num*10 + int(ch-'0')
// 		}else {
// 			return
// 		}
// 	}

// 	result := 0
// 	for i := 1; i <= 9; i ++ {
// 		result = i*num

// 		z01.PrintRune(rune(i + '0'))
// 	    z01.PrintRune(' ')
// 		z01.PrintRune('x')
// 		z01.PrintRune(' ')
// 		Itoa(num)
// 		z01.PrintRune(' ')
// 		z01.PrintRune('=')
// 		z01.PrintRune(' ')
// 		Itoa(result)
// 		z01.PrintRune('\n')
// 	}
// }

//	func Itoa(num int){
//		str := ""
//		for num != 0 {
//			mod := num % 10
//			rune1 := '0'
//			for i:= 0; i < mod; i++{
//				rune1 ++
//			}
//			str = string(rune1) + str
//			num = num / 10
//		}
//		for _, chr := range str{
//			z01.PrintRune(chr)
//		}
//	}
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