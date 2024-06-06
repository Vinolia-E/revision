// package main

// import "fmt"

//	func main(){
//		num := 123
//		result := ""
//
//	for num != 0{
//		mod := num % 10
//		startrune := '0'
//		for i := 0; i < mod; i++ {
//			startrune++
//		}
//		result = string(startrune) + result
//		num = num/10
//	}
//
// fmt.Println (result)
// }
package main

import "fmt"

func main() {
	num := 123
	result := ""
	for num != 0 {
		mod := num%10
		run := '0'
		for i := 0; i < mod; i++ {
			run++
		}
		result = string(run) + result
		num = num/10
	}
	fmt.Println(result)
}

