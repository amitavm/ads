package main

/*
This is an implementation of the heap-sort algorithm using max heaps.  A max
heap is a binary-tree like structure in which each inner (parent) node has a
value that is not less than the values of its children.  We simulate such a
"tree" using arrays (slices): see the parent(), lchild() and rchild() functions
below.

Refer to any good textbook on ADS to learn about (or review) this algorithm.
If you go through CLRS, you will notice the similarity of our approach with
that book---including the names of the functions!
*/

func main() {
	testSort(heapSort)
}

// Sort the elements in arr in ascending order using the heap-sort algorithm.
func heapSort(arr []int) {
	for i := len(arr); i > 1; i-- {
		buildMaxHeap(arr[:i])
		arr[0],arr[i-1] = arr[i-1],arr[0]
	}
}

// Turn an array of (random) integers into a max-heap.
func buildMaxHeap(arr []int) {
	// Starting with the last parent, traverse all the (parent) nodes backwards
	// and ensure that the max-heap property holds for all of them.  (Q: Why do
	// the traversing backwards?)
	for i := parent(len(arr)-1); i >= 0; i-- {
		maxHeapify(arr, i)
	}
}

// Ensure that the max-heap property holds at index i in arr.
// All other positions of arr are assumed to uphold the max-heap property.
func maxHeapify(arr []int, i int) {
	m := i // Index of the max of the three nodes.
	l := lchild(i)
	if l < len(arr) && arr[l] > arr[m] {
		m = l
	}
	r := rchild(i)
	if r < len(arr) && arr[r] > arr[m] {
		m = r
	}
	if m != i {
		arr[i], arr[m] = arr[m], arr[i]
		maxHeapify(arr, m)
	}
}

// Return the index of the parent node of node i.
func parent(i int) int {
	return (i - 1) / 2
}

// Return the index of the left-child of node i.
func lchild(i int) int {
	return 2*i + 1
}

// Return the index of the right-child of node i.
func rchild(i int) int {
	return 2*i + 2
}
