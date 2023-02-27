# Sorting Algorithms

This directory contains implementations of some classic sorting algorithms in
Go.  We have included *brief* descriptions for the algorithms in the sections
below.

However, you should treat this is only as a reference.  If you are learning
these sorting algorithms, it's much better to learn them from a good algorithms
textbook.

## Running the Algorithms

If you have Go installed on your system, you can run these algorithms using a
command line similar to this:

```
$ go run insertion-sort.go utils.go
```

(This shows how you would run the insertion-sort algorithm.  You just need to
use the correct source file name for the other algorithms.)

This shows the command line for a Unix-like system (indicated by the leading
`$`, which conventionally represents a Bourne command shell), including Mac
systems.  But the same command line should work on a Windows command prompt
as well.

## Insertion Sort

[Here](./insertion-sort.go) is an implementation of the classic insertion sort
algorithm in Go.  The algorithm works similar to how people normally sort a
hand of cards: consider each card starting with the second left-most, and put
it into its "right place" in the group of cards to its left.  The cards to the
left of the one under consideration is assumed to be already in sorted
(ascending) order, of course.

Even though the runtime complexity of insertion-sort is $O(n^2),$ it's supposed
to be one of the fastest quadratic-time sorting algorithms.  Its advantages
are:

* It is relatively short, and simple to understand.
* It doesn't require any special pre- or post-processing.
* It is an in-place sorting algorithm, so it doesn't require extra storage.

Due to these reasons, the constant factor $c$ in its runtime complexity tends
to be small.  So it's often used in other algorithms&mdash;where sorting
appears as a sub-task&mdash;instead of other, more sophisticated sorting
algorithms when the arrays to be sorted are known to be not too big.

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

## Heap Sort

Heap sort is an unusual and clever sorting algorithm.  Unlike the other sorting
algorithms we discuss here, this one is not that easy to "guess".  Insertion
sort is a straight-forward translation into code of a process many of us are
familiar with: as we mentioned above, it's basically the procedure we normally
follow to sort a hand of cards.  Merge sort almost "suggests itself" once you
get the idea of the divide-and-conquer approach.  Even quick-sort is
essentially a divide-and-conquer style algorithm, with a subtle improvement
over merge-sort, of course.

However, coming up with heap-sort will likely require a genius, in a manner of
speaking!  It probably came up in the context of some research work.  To sort
an array, which is a *linear* structure, we simulate a "tree" (a *nonlinear*
structure), called a *max heap*, right in the array, and impose a
constraint---the max-heap constraint---on its elements repeatedly to come up
with the "next biggest element".  That is the gist of it; you can see the
details in the implementation.

Apart from its cleverness, another nice aspect of heap-sort is its runtime
complexity: it runs in $O(n\lg n)$ time!  Even though quick-sort is mostly the
sorting algorithm of choice in practice, heap-sort does have some good things
to recommend it:

* It has one of the fastest runtimes among sorting algorithms.
* It's an in-place sorting algorithm.
* It gives a good amount of intellectual satisfaction when understood!

[Here](./heap-sort.go) is one way to implement heap-sort.  You will probably
need to go through an ADS textbook to understand it first, before the code
makes much sense.
