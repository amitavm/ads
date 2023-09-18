// Generate all possible permutations of the elements in a given array/slice.
// Here we permute a slice of int's to keep things simple.  But the same idea
// should apply to a slice of anything.

package main

import (
	"fmt"
)

func permute(arr []int) [][]int {
	return aux(arr, []int{})
}

func aux(arr, perm []int) [][]int {
	if len(arr) == 0 {
		return [][]int{perm}
	}

	result := [][]int{}
	for i, e := range arr {
		newarr := []int{}
		newarr = append(newarr, arr[:i]...)
		newarr = append(newarr, arr[i+1:]...)

		newperm := []int{}
		newperm = append(newperm, perm...)
		newperm = append(newperm, e)

		result = append(result, aux(newarr, newperm)...)
	}
	return result
}

func main() {
	arr := []int{1, 2, 3}
	perms := permute(arr)
	fmt.Printf("found %d permutations:\n", len(perms))
	for _, p := range perms {
		fmt.Printf("    %v\n", p)
	}
}
