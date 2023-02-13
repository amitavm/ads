# Sorting Algorithms

This directory contains some classic sorting algorithms:

## Insertion Sort

[Here](./insertion-sort.go) is an implementation of the classic insertion sort
algorithm.  Even though it's runtime complexity is $O(n^2)$, it is supposed to
be one of the fastest quadratic-time sorting algorithms.  It is often a good
choice when the array to sort is known to be not too big.  It's advantages are:

* It is relatively short, and simple to understand.
* It doesn't require any special pre-processing.
* It is an in-place sorting algorithm, so it doesn't require extra storage.

So the constant factor $c$ in its runtime complexity tends to be small.  For
these reasons, insertion-sort is sometimes used in other, more extensive
algorithms where sorting relatively small arrays appears as a sub-task.

<!-- TODO: Add an example? -->

## Merge Sort

* [merge sort](./merge-sort.go)
