package main 

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	args := os.Args[1]
	if len(os.Args) != 2{
		return
	}

	if len(args) < 1 {
		fmt.Println(0)
	}

	//n := args[1]
	num, _ :=  strconv.Atoi(args)

	power2 := Power(num)
	fmt.Println(power2)
	
}

func Power(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n-1)) == 0
}
