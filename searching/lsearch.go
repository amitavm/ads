// An implementation of the linear-search algorithm.

package main

import (
	"fmt"
)

// Search for `val` in `arr` using the linear-search algorithm.
// Return its index if found, else return -1.
func lsearch(arr []int, val int) int {
	for i, e := range arr {
		if e == val {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{2, -3, 44, 59, 17, -40, 37}
	val := 3
	idx := lsearch(arr, val)
	if idx == -1 {
		fmt.Printf("%d does NOT occur in arr\n", val)
	} else {
		fmt.Printf("%d occurs at index %d in arr\n", val, idx)
	}
}
