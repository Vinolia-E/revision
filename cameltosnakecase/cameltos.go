/*Instructions

Write a function that converts a string from camelCase to snake_case.

    If the string is empty, return an empty string.
    If the string is not camelCase, return the string unchanged.
    If the string is camelCase, return the snake_case version of the string.

For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

    lowerCamelCase
    UpperCamelCase

Rules for writing in camelCase:

    The word does not end on a capitalized letter (CamelCasE).
    No two capitalized letters shall follow directly each other (CamelCAse).
    Numbers or punctuation are not allowed in the word anywhere (camelCase1).

Expected function

func CamelToSnakeCase(s string) string{

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

And its output:

$ go run .
Hello_World
hello_World
camel_Case
CAMELtoSnackCASE
camel_To_Snake_Case
hey2
*/

package main

import "fmt"

func main() {
	//
	fmt.Println(CamelToSnakeCase("HelloThereWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelLCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
	// fmt.Println(Validatecamel("HD"))
	// fmt.Println(Validatecamel("helloWorld"))
	// fmt.Println(Validatecamel("camelCase"))
	// fmt.Println(Validatecamel("CAMELtoSnackCaSE"))
	// fmt.Println(Validatecamel("camelToSnakeCase"))
	// fmt.Println(Validatecamel("hey2"))
}

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := ""
	//final := ""
	if Validatecamel(s) {
		// for i, ch := range s {
		// 	if IsUpper(ch) && i != 0 {
		// 		result = result + "_" + string(ch)
		// 		//final += result
		// 		//result = ""
		// 	} else {
		// 		result += string(ch)
		// 	}
		// }
		// for i := 0; i < len(s); i++ {
		// 	if IsUpper(rune(s[i])) && i != 0 {
		// 		result = s[:i] + "_"
		// 		final += result
		// 		result = ""
		// 	}
		// 	// i++
		// }
		for i , v := range s {
			if IsUpper(v) && i != 0 {
				if !IsUpper(rune(s[i+1])) { 
					result = result + "_" + string(v)
				} else {
					result = result + string(v)
				}
			}
		}
	} else {
		return s
	}
	return result
}

func Validatecamel(s string) bool {
	for i, ch := range s {
		if !IsUpper(ch) && !IsLower(ch) {
			return false
		}
		if i == len(s)-1 && IsUpper(ch) {
			return false
		}
		if IsUpper(ch) && IsUpper(rune(s[i+1])) {
			return false
		}
		if i == len(s)-1 && IsUpper(ch) {
			return false
		}
	}
	return true
}

func IsUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func IsLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}
