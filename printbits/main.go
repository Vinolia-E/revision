package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	first, err := strconv.Atoi(args)
	num := strconv.FormatInt(int64(first), 2)
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Println(num)
	
}