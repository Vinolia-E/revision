// Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.

//     The display will be followed by a newline ('\n').

//     If the number of arguments is different from 2, the program displays nothing.

// Usage

// $ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
// padinto
// $ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
// df6ewg4
// $

// package main

// import (
// 	//"fmt"
// 	"os"
// 	"github.com/01-edu/z01"
// )

// func main() {
    //if len(os.Args) != 3 {
        // 		return
        // 	}
// 	args := os.Args[1:]
// 	s1 := args[0]
// 	s2 := args[1]
// 	result := ""
// 	interMap := make(map[rune]bool)


// 	for _, char := range s2 {
// 		interMap[char] = true
// 	}

// 	for _, char := range s1 {
// 		if interMap[char] {
// 			result += string(char)
// 			interMap[char] = false
// 			//delete(interMap, char)
// 		}
// 	}
// 		for _, chr := range result {
// 			z01.PrintRune(chr)
// 		} 
	
// 	z01.PrintRune('\n')
// }



package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        return
    }
    str1 := os.Args[1]
    str2 := os.Args[2]
    result := ""
    seen := make(map[rune]bool)
    for _, c := range str1 {
        if contains(str2, c) && !seen[c] {
            result += string(c)
            seen[c] = true
        }
    }
    fmt.Println(result)
}

func contains(s string, r rune) bool {
    for _, c := range s {
        if c == r {
            return true
        }
    }
    return false
}
