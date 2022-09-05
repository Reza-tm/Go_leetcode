package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := fuzzBuzz(50000)
	fmt.Println(arr)
}

func fuzzBuzz(len int) []string {
	tmp := []string{}
	for i := 1; i <= len; i++ {
		switch j := i; {
		case j%3 == 0 && j%5 == 0:
			tmp = append(tmp, "FizzBuzz")
			continue
		case j%3 == 0:
			tmp = append(tmp, "Fizz")
		case j%5 == 0:
			tmp = append(tmp, "Buzz")
		default:
			tmp = append(tmp, strconv.Itoa(i))
		}

	}
	return tmp
}
