package max

func Max(x []int) int {
	if len(x) == 0 {
		return 0
	}
	for i := 0 ; i < len(x); i++{
		if x[i] > x[0]{
			x[0]=x[i]
		}
	}
	return x[0]
}
