package main

import "fmt"

func main() {
	exArr := []int{3, 1, 2, 10, 1}
	runningSum(exArr)
}

func runningSum(nums []int) []int {
	tmp := []int{}
	for idx, _ := range nums {
		sum := 0
		for i := 0; i <= idx; i++ {
			fmt.Println(nums[i])
			sum += nums[i]
		}
		tmp = append(tmp, sum)
	}
	return tmp
}
