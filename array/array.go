package array

import "fmt"

func RemoveDuplicados(arr []int) []int {
	result := []int{}
	for _, v := range arr {
		if !exists(result, v) {
			result = append(result, v)
		}
	}

	fmt.Errorf("erro %d", 123)

	return result
}

func exists(arr []int, v int) bool {
	for i := range arr {
		if v == arr[i] {
			return true
		}
	}

	return false
}
