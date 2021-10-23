package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := 0
	b := 1
	i := 0
	return func() int {
		if i == 0 {
			i++
			return 0
		} else if i == 1 {
			i++
			return 1
		} else {
			b += a
			a = b - a
			return b
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
