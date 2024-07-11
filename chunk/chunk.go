package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)

}

func Chunk(slice []int, size int) {
	// if size <= 0 {
	// 	fmt.Println()
	// 	return
	// }
	// // if len(slice) == 0 {
	// // 	fmt.Println([]int64{})
	// // 	return
	// // }
	// var chunk  [][]int
	// for i := 0; i < len(slice); i+=(size) {
		
	// 	split := i + size
	// 	if split > len(slice) {
	// 		split = len(slice)
	// 	}
	// 	sub := slice[i:split]
	// 	chunk = append(chunk, sub)
	// }
	// fmt.Println(chunk)

var chunk [][]int
if size == 0 {
	fmt.Println()
	return
}

if len(slice) == 0 {
	fmt.Println([]int{})
	return
}

	for i := 0; i <= len(slice); i+=size {
	//	chunks := []int{}
		div := i+size
		if div > len(slice) {
			div = len(slice)
		}
		chunks := slice[i:div]
		chunk = append(chunk, chunks)
	}
	fmt.Println(chunk)
}