package main

import (
	"fmt"
	"myrevision/itoa"
	"myrevision/max"
	//"myrevision/itoa"
	//"os"
)

func main() {
	a := []int{7881, 327, 34, 8, 54, 36, 655}
	max :=max.Max(a)
	fmt.Println(max)
	itoa  := itoa.Itoa(12345)
	fmt.Println("Itoa", itoa)
	fmt.Println(atoi)

	//fmt.Println(CountNegative([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
    //fmt.Println(CountNegative([]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}))
    //fmt.Println(CountNegative([]int{}))
    fmt.Println(CountNegative([]int{-1,2,0,-3,-5}))

}


func CountNegative(numbers []int) int {
	count := 0
	if len(numbers) == 0 {
		return 0
	}
	for _,num := range numbers{
		if num < 0 {
			count++
		}
	}
	return count

	
}

