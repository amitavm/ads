// An implementation of the binary-search algorithm.

package main

import (
	"fmt"
	"slices"
)

// Search for `val` in `arr` using the binary-search algorithm.
// Return its index if found, else return -1.
// (We *assume* that the input array is sorted.)
func bsearch(arr []int, val int) int {
	lo, hi := 0, len(arr)
	for lo <= hi {
		mid := lo + (hi - lo)/2
		if arr[mid] == val {
			return mid
		}
		else if arr[mid] < val {
			lo = mid+1
		}
		else {
			hi = mid-1
		}
	}
	return -1
}

func main() {
	arr := []int{2, -3, 44, 59, 17, -40, 37}
	fmt.Println(arr)
	slices.Sort(arr)
	fmt.Println(arr)

	val := 3
	idx := bsearch(arr, val)
	if idx == -1 {
		fmt.Printf("%d does NOT occur in arr\n", val)
	} else {
		fmt.Printf("%d occurs at index %d in arr\n", val, idx)
	}
}
