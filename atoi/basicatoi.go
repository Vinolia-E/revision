package atoi

//import "fmt"

func Batoi(s string) int {
	num := 0
	for _, ch := range s {
		num = num*10 + int(ch + '0')
	}
	return num
}