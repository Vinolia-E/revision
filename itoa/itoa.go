package itoa

func Itoa(x int) string {
	var intSlc []int
	//var newSlc []int
	var div int
	var mod int
	var str string

	div = x
	for div > 0 {
		mod = div % 10
		div = div / 10

		intSlc = append(intSlc, mod)
	}

	for i := len(intSlc) - 1; i >= 0; i-- {
		str = str + string(int32(intSlc[i]+48))
	}

	// for i := 0; i < len(newSlc); i++ {
	// 	str = str + string(rune(newSlc[i]+48))
	// }
	return str
}
