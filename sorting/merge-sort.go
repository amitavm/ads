package main

func main() {
	testSort(mergeSort)
}

// Sort an array of (random) integers using the merge-sort algorithm.
// (The elements are sorted in ascending order.)
func mergeSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return // Nothing to do; arr is already sorted (trivially).
	}

	m := n / 2 // Index of the "middle" element.
	mergeSort(arr[:m])
	mergeSort(arr[m:])
	merge(arr, m, n)
}

// Merge the elements in the subarrays arr[0:m] and arr[m:e].  The elements in
// the subarrays are assumed to be already in sorted (ascending) order.
func merge(arr []int, m int, e int) {
	res := make([]int, e) // Temporary work area where we merge the elements.
	var i, j, k int       // These need to be visible beyond the first loop.

	// Copy elements as long as both subarrays have elements to consider.
	for i, j, k = 0, m, 0; i < m && j < e; k++ {
		if arr[i] < arr[j] {
			res[k] = arr[i]
			i++
		} else {
			res[k] = arr[j]
			j++
		}
	}

	// Copy left-over elements, if any.
	// At most only one of these loops will execute.
	for ; i < m; i, k = i+1, k+1 {
		res[k] = arr[i]
	}
	for ; j < e; j, k = j+1, k+1 {
		res[k] = arr[j]
	}

	// Remember to copy the merged elements back into arr before returning.
	copy(arr, res)
}
