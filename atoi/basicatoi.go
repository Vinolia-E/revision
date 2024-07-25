package main

import "fmt"

// import "fmt"
func Batoi(s string) int {
	num := 0
	for _, ch := range s {
		num = num*10 + int(ch-'0')
	}
	return num
}

func Atoi2(s string) int {
	mult := 1
	num := 0
	for i, ch := range s {
		if i == 0 && ch == '-' {
			mult = -1
			continue
		} else if i == 0 && ch == '+' {
			mult = 1
			continue
		}
		if ch < '0' && ch > '9' {
			mult = 0	
		}
		num = num*10 + int(ch-'0')
	}
	return num * mult
}

func main() {
	fmt.Println(Atoi2("-123"))
	fmt.Println(Batoi("23"))
	x := Atoi2("-123")
	y := Batoi("23")
	fmt.Println(x + y)
}
