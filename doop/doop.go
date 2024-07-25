package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	s1 := args[0]
	operator := args[1]
	s2 := args[2]
	var calculation int
	answer := ""

	if operator == "%" && s2 == "0" {
		fmt.Println("No modulo by 0")
		return
	}else if operator == "/" && s2 == "0" {
		fmt.Println("No division by 0")
		return
	}else if operator == "+" {
		calculation = Atoi(s1) + Atoi(s2)
		answer = Itoa(calculation) 
	} else if operator == "-" {
		calculation = Atoi(s1)- Atoi(s2)
		answer = Itoa(calculation)
	} else if operator == "%"{
		calculation = Atoi(s1)%Atoi(s2)
		answer = Itoa(calculation)
	} else if operator == "*" {
		calculation = (Atoi(s1)) * (Atoi(s2))
		answer = Itoa(calculation)
	} else if operator == "/" {
		calculation = Atoi(s1)/ Atoi(s2)
		answer = Itoa(calculation)
	}else {
		return
	}
	//fmt.Println(calculation)
	 
	// if calculation < -9223372036854775807 {
	// 	fmt.Println("there")
	// 	return
	// }
	// if uint(calculation) > 9223372036854775807 {
	// 	fmt.Println("Heren")
	// 	return
	// }
	fmt.Println(answer)
}

func Atoi(s string) int {
	result := 1
	var num uint
	if s[0] == '-' {
		result = -1
		s = s[1:]
	}
	if s[0] == '+' {
		s = s[1:]
	}
	for _, ch := range s {
		// if ch == '-' && i == 0 {
		// 	result = -1
		// 	continue
		// } else if ch == '+' && i == 0 {
		// 	result = 1
		// 	continue
		// }
		if ch < '0' || ch > '9' {
			os.Exit(0)
		}
		num = num*10 + uint(ch-'0')
	}
	 if num > 9223372036854775807{
	 	fmt.Println("Overflow")
	 	os.Exit(0)
	 }
	 //fmt.Println(int(num) * result)
	return int(num) * result
}

func Itoa(n int) string {
	mod := 0
	result := ""
	if n < 0 {
		n = n*-1
		for n > 0 {
			mod = n%10
			result = string(rune(mod)+'0')+result
			n = n/10
		}
		return "-"+result
	}
	if n == 0 {
		return "0"
	}

	for n > 0 {
		mod = n%10
		result = string(rune(mod)+'0')+result
		n = n/10
	}
	return result
}

