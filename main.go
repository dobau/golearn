package main

import (
	"fmt"
)

func main() {
	var x int = 1
	y := 2

	fmt.Println("TESTE")
	fmt.Printf("x + y = %d\n", (x + y))

	var even = make([]int, 0, 10)
	odd := []int{}
	for i := 0; i < 10; i += 1 {
		if i%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}

	}

	fmt.Printf("even = %v\n", even)
	fmt.Printf("even = %+v\n", odd)
}
