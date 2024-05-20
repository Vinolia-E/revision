package max

import (
)

func Max(x []int) int {
	if len(x) == 0 {
		return 0
	}
	max := x[0]
	for _, number := range x {
		if number > max {
			max = number
		}
	}
	return max
}