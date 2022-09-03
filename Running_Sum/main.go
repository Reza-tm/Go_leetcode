package main

import "fmt"

func main() {
	exArr := []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum(exArr))
}

func runningSum(nums []int) []int {
	var tmp = make([]int, len(nums))
	tmp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		tmp[i] = nums[i] + tmp[i-1]
	}
	return tmp
}
