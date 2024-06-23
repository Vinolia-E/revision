package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}

	str1 := args[0]
	str2 := args[1]
	count := 0
	result := ""

	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2) && count < len(str1); j++ {
			if str1[i] == str2[j] {
				for k := range str2 {
					if str2[j] != str2[k] && k != j {
						result += string(str1[i])
						count++
					} else {
						count++
					}
				}
				if count == len(str1)-1 {
					fmt.Println(result)
					return
				}
			}
		}
	}

}

// package main

// import (
//     "fmt"
//     "os"
// )

// func main() {
//     if len(os.Args) != 3 {
//         return
//     }
//     str1 := os.Args[1]
//     str2 := os.Args[2]
//     result := ""
//     seen := make(map[rune]bool)
//     for _, c := range str1 {
//         if contains(str2, c) && !seen[c] {
//             result += string(c)
//             seen[c] = true
//         }
//     }
//     fmt.Println(result)
// }

// func contains(s string, r rune) bool {
//     for _, c := range s {
//         if c == r {
//             return true
//         }
//     }
//     return false
// }
