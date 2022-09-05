package main

import "fmt"

func main() {
	fmt.Println(stepCounter(14))
}

func stepCounter(n int) int {
	tmp := n
	steps := 0

	for {
		if tmp == 0 {
			break
		}
		if tmp%2 == 0 {
			tmp /= 2
			steps++
		} else {
			tmp--
			steps++
		}
	}

	return steps
}
