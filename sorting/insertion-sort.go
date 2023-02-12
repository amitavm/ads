package main

import (
	"fmt"
)

// insertionSort sorts an array of integers into ascending order in place,
// using the insertion-sort algorithm.
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j, key := i-1, arr[i]
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{-2, 44, 36, -19, 0, 58, 7, -69}
	fmt.Println(arr)
	insertionSort(arr)
	fmt.Println(arr)
}
