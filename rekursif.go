package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Recursive Binary Search
func recursiveBinarySearch(arr []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return recursiveBinarySearch(arr, mid+1, high, target)
	} else {
		return recursiveBinarySearch(arr, low, mid-1, target)
	}
}

// Iterative Binary Search
func iterativeBinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Generate Random Array
func generateArray(size int, sortedArray bool) []int {
	array := make([]int, size)
	for i := range array {
		array[i] = rand.Intn(1000)
	}
	if sortedArray {
		sort.Ints(array)
	}
	return array
}

// Measure Execution Time for Recursive
func measureRecursiveTime(array []int, target int, sorted bool) float64 {
	if !sorted {
		sort.Ints(array)
	}
	start := time.Now()
	recursiveBinarySearch(array, 0, len(array)-1, target)
	return time.Since(start).Seconds()
}

// Measure Execution Time for Iterative
func measureIterativeTime(array []int, target int, sorted bool) float64 {
	if !sorted {
		sort.Ints(array)
	}
	start := time.Now()
	iterativeBinarySearch(array, target)
	return time.Since(start).Seconds()
}

// Main Program
func main() {
	sizes := []int{100, 500, 1000, 5000, 10000}
	fmt.Printf("%-10s %-20s %-20s %-20s %-20s\n", "Size", "Recursive Sorted", "Recursive Unsorted", "Iterative Sorted", "Iterative Unsorted")

	for _, size := range sizes {
		array := generateArray(size, false)
		target := array[rand.Intn(len(array))]

		// Measure Recursive Times
		recursiveSortedTime := measureRecursiveTime(generateArray(size, true), target, true)
		recursiveUnsortedTime := measureRecursiveTime(array, target, false)

		// Measure Iterative Times
		iterativeSortedTime := measureIterativeTime(generateArray(size, true), target, true)
		iterativeUnsortedTime := measureIterativeTime(array, target, false)

		// Display Results
		fmt.Printf("%-10d %-20.10f %-20.10f %-20.10f %-20.10f\n",
			size, recursiveSortedTime, recursiveUnsortedTime, iterativeSortedTime, iterativeUnsortedTime)
	}
}
