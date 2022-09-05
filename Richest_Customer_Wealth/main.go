package main

import "fmt"

func main() {
	exArr := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println("answer is : ", richFinder(exArr))
}

func richFinder(arr [][]int) int {
	var tmp = make([]int, len(arr))
	sumTmp := 0
	max := 0
	for idx, val := range arr {
		for _, k := range val {
			sumTmp += k
		}
		tmp[idx] = sumTmp
		sumTmp = 0
	}
	for _, val := range tmp {
		if val > max {
			max = val
		}
	}
	return max
}
