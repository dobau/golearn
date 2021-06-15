package main

import (
	"fmt"

	"github.com/dobau/golearn/outropacote"
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

	outropacote.FuncaoPublica()

	Loop1()
}

func Loop1() {

	slice1 := []string{"a", "b"}
	slice2 := make([]string, 2)
	array1 := [2]string{"a", "b"}
	array2 := [2]string{"a", "b"}

	slice2[0], slice2[1] = "a", "b"

	for i := range slice1 {
		fmt.Printf("[%d] slice1(%s) slice2(%s) array1(%s) array2(%s)\n",
			i, slice1[i], slice2[i], array1[i], array2[i])
	}
}
