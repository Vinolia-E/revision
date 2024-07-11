package main 

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		return
	}

	if len(args) == 1 {
		fmt.Println(0)
	}

	//n := args[1]

	
}
