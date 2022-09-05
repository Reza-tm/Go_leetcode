package main

import (
	"math"
)

func main() {
	middleNode([]int{1, 2, 3, 4})
}

func middleNode(head []int) []int {
	arrLen := float64(len(head))
	middle := math.Ceil(float64(arrLen) / 2.0)
	tmp := []int{}
	for i := middle + 1; i <= arrLen; i++ {
		tmp = append(tmp, int(i))
	}
	return tmp
}
