package main

import "fmt"

func main() {
	var arr = []int{4, 5, 2, 7, 1, 6, 3}
	fmt.Println("Before bubble-sort: ", arr)
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	print("-----------------------------------------------\n")
	fmt.Println("After bubble-sort: ", arr)
}
