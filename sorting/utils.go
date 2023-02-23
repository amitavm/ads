package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Return a slice populated with n random integers in [min, max).
func randArr(n int, min int, max int) []int {
	res := []int{}  // The result slice to return.
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(max-min)+min)
	}
	return res
}

// Check if a given array (slice) of integers is sorted (in ascending order).
func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// Test drive a given sort function.
func testSort(sort func([]int)) {
	arr := randArr(15, -50, 50)
	fmt.Println(arr)
	sort(arr)
	fmt.Println(arr)
	fmt.Printf("arr is sorted: %v\n", isSorted(arr))
}
