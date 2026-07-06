package main

import (
	"fmt"
)

func sliceArr(scores []int) []int {
	length := len(scores)
	mid := length / 2 // get the half of array length
	left := []int{}
	right := []int{}

	left = scores[0:mid]
	right = scores[mid+1 : length]

	left = append(left, 88) // append new data

	scores = append(left, right...) // reassign with new arrays
	return scores
}

func main() {
	scores := sliceArr([]int{50, 75, 66, 20, 22, 32, 90})
	for x := range len(scores) {
		fmt.Println(scores[x])
	}
}
