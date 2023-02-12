package main

import (
	"fmt"
)

// mergeSort sorts an array of integers into ascending order in place, using
// the merge-sort algorithm.
func mergeSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return // Nothing to do, it's trivially sorted.
	}

	p := n / 2
	mergeSort(arr[:p])
	mergeSort(arr[p:])
	merge(arr)
}

// Merge the two sorted halves of arr into one sorted sequence.
func merge(arr []int) {
	n := len(arr)
	p := n / 2

	res := make([]int, n)
	i, j, k := 0, p, 0
	for ; i < p && j < n; k++ {
		if arr[i] < arr[j] {
			res[k] = arr[i]
			i++
		} else {
			res[k] = arr[j]
			j++
		}
	}

	for ; i < p; k++ {
		res[k] = arr[i]
		i++
	}

	for ; j < n; k++ {
		res[k] = arr[j]
		j++
	}

	copy(arr, res)
}

func main() {
	arr := []int{-2, 44, 36, -19, 0, 58, 7, -69}
	fmt.Println(arr)
	mergeSort(arr)
	fmt.Println(arr)
}
