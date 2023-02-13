# Sorting Algorithms

This directory contains some classic sorting algorithms, each of which is
described in some detail in the corresponding sections below.

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

[Merge sort](./merge-sort.go) is one of the standard examples of
divide-and-conquer algorithms.  It's runtime complexity is $O(n\log n)$, but it
is rarely used in practice.  Its biggest downside is probably the fact that it
uses extra storage in its "merge" step (which is where its name comes from).
Meaning, it's not an in-place sorting algorithm.  Still, it's worth learning
because of two reasons:

* Its a good example to illustrate the divide-and-conquer approach in algorithm
  design.
* The merge procedure used in this algorithm should be learned/known by every
  "student" of the field of algorithms (and data structures).
