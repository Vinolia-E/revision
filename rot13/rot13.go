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
	fmt.Println(Rot14("HEY Alice you re 12345 Z"))
	fmt.Println(Rot13("a Z n N"))
	fmt.Println(Rot13("HEY Alice you re 12345 Z"))
	fmt.Println(Rot10("HEY Alice you re 12345 Z"))
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

func Rot10(s string) string {
	result := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			ch = (ch - 'a' + 10) %26 + 'a'
			result += string(ch)
		} else if ch >= 'A' && ch <= 'Z' {
			ch = (ch - 'A' + 10)% 26 + 'A'
			result += string(ch)
		} else {
			result += string(ch)
		}
	}
	return result
}