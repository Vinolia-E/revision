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
	if size == 0 {
		fmt.Println()
		return
	}
	if len(slice) == 0 {
		fmt.Println([]int{})
		return
	}

	var result [][]int
	var chunk []int
	for len(slice) >= size {
		chunk, slice = slice[:size], slice[size:]
		result = append(result, chunk)
	}
	if len(slice) > 0 {
		result = append(result, slice[:])
	}
	fmt.Println(result)
}