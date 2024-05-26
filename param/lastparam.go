package main

import (
	//"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	 //arg := args[0]
	//arg := args[len(args)-1]

	for _, word := range args {
		//z01.PrintRune(word)
		for _, char := range word {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
	
	// z01.PrintRune('\n')

	// count := 0

	// for _, x := range args {
	// 	for range x {
	// 		count++
	// 	}
	// }
	//fmt.Println(count)
	// str := ""
	// startrune := '0'

	// if count != 0 {
	// 	mod := count%10
	// 	div := count/10
		
	// 	for i := 0; i < mod; i++ {
	// 		startrune++
	// 	}
	// 	str += string(startrune)
	// 	count = div
	// }
	// for _, ch := range str {
	// 	z01.PrintRune(ch)
	// }
	// z01.PrintRune('\n')

	//args := os.Args[1:]
	count := len(args)

	var result []rune

	if len(args) == 0 {
		//return 
		z01.PrintRune('0')
	}
	for count > 0 {
		digit := count % 10
		result = append([]rune{rune('0' + digit)}, result...)
		count /= 10
	}
	for _, ch := range result {
        z01.PrintRune(ch)
    }
	 z01.PrintRune('\n')	
}	
