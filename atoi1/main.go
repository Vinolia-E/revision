package main

import "fmt"

func Atoi(str string) int {
	// multiplier := 1
	// if str[0] == '-' {
	// 	multiplier = -1
	// 	str = str[1:]
	// } else if str[0] == '+'{
	// 	multiplier = 1
	// 	str = str[1:]
	// }

	// result := 0

	// for _, ch := range str{
	// 	result = result * 10 + int(ch - '0')
	// }
	// return result * multiplier
	multiplier := 1
	result := 0

	for i, ch := range str {
		if i == 0 && ch == '-' {
			multiplier = -1
			continue
		} else if i == 0 && ch == '+' {
			multiplier = 1
			continue
		}
		if ch < '0' || ch > '9' {
			return 0
		} else {
			result = result*10 + int(ch - '0')
		}
	}
	return result * multiplier
}

func main() {
	fmt.Println(Atoi("-254"))
	fmt.Println(+543)
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
