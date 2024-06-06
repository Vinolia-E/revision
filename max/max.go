package max

func Max(x []int) int {
	max := x[0]
	if len(x) == 0 {
		return 0
	}
	for i := 0 ; i < len(x); i++{
		if x[i] > max{
			max=x[i]
		}
	}
	return max
}
