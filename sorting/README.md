# Sorting Algorithms

This directory contains implementations of some classic sorting algorithms in
Go.  We have included *brief* descriptions for the algorithms in the sections
below.

However, you should treat this is only as a reference.  If you are learning
these sorting algorithms, it's much better to learn them from a good algorithms
textbook.

## Insertion Sort

[Here](./insertion-sort.go) is an implementation of the classic insertion sort
algorithm in Go.  Even though its runtime complexity is $O(n^2)$, it's supposed
to be one of the fastest quadratic-time sorting algorithms.  It is often a good
choice when the array to sort is known to be not too big.  It's advantages are:

* It is relatively short, and simple to understand.
* It doesn't require any special pre-processing.
* It is an in-place sorting algorithm, so it doesn't require extra storage.

So the constant factor $c$ in its runtime complexity tends to be small.
Therefore, insertion-sort is sometimes used in other algorithms where sorting
relatively small arrays appears as a sub-task.

<!-- TODO: Add an example? -->

## Merge Sort

[Merge sort](./merge-sort.go) is one of the standard examples of
divide-and-conquer algorithms.  Its runtime complexity is $O(n\log n)$, which
looks much better than quadratic-time algorithms like insertion sort, at least
on paper.

However, it is rarely used in practice because there are other algorithms with
similar run times that perform better.  (Quick-sort, for example.)  Its biggest
downside is probably the fact that it uses extra storage in its "merge" step
(which is where its name comes from).  Meaning, it's not an in-place sorting
algorithm.  Still, it's worth learning because of two reasons:

* It is a good example to illustrate the divide-and-conquer approach in
  algorithm design.
* The merge procedure used in this algorithm should be learned/known by every
  "student" of algorithms (and data structures)!
