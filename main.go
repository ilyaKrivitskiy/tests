package main

import "fmt"

var p = fmt.Printf

func main() {
	var arr = []int{}
	arr = []int{2, 5, 3, 6, 1, 7, 9, 4, 8}
	var res int = 0
	for i, value := range arr {
		res += value * (i + 1)
	}
	p("Result is %d\n", res)
}
