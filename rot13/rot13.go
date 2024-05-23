package main

import "fmt"

func Rot13(str string)string{
	new := []rune{}
	result := ""
	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' {
			new = []rune{(ch - 'a' + 13) % 26 + 'a'}
			result = result + string(new)
		} else if ch >= 'A' && ch <= 'Z' {
			new = []rune{(ch - 'A' + 13) % 26 + 'A'}
			result = result + string(new)
		} else {
			result= result + string(ch)
		}
	}
	return result
}

func main(){
	fmt.Println(Rot13("HEY Alice you re 12345 Z"))
	fmt.Println(Rot14("z Z n N"))
}

func Rot14(s string) string{
	replace := []rune{}
	output := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z'{
			replace = []rune{(ch - 'a' +14)%26 + 'a'}
			output += string(replace)
		} else if ch >= 'A' && ch <= 'Z' {
			replace = []rune{(ch-'A'+14)%26 + 'A'}
			output += string(replace)
		} else {
			output += string(ch)
		}
	}
	return output
}