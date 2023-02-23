package main

func main() {
	testSort(insertionSortV2)
}

// Sort an array of integers in-place, using the insertion-sort algorithm.
// This is a somewhat simplistic implementation of insertion-sort; we will
// improve on this a little in the next version below.
func insertionSortV1(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			// Keep shifting the element under consideration (arr[i]) to the
			// left as long as it's less than the one to its left.
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

// The above version looks ok (and it's expressed concisely in Go), but we can
// improve on it a bit: instead of swapping the elements, we can keep moving
// them to the right till we find an element that is <= the element being
// considered (the "key").  So we effectively do only one copy instead of two
// in each iteration (of the inner loop).  Ref: CLRS4:19
func insertionSortV2(arr []int) {
	for i := 1; i < len(arr); i++ {
		j, key := i-1, arr[i]
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j] // Move elements > key one position to the right.
		}
		arr[j+1] = key
	}
}
