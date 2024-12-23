package main

import (
	"fmt"
	"math/rand"ie
	"sort"
	"time"
)

// Binary Search Function
func binarySearch(arr []int, target int) int {
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

// Measure Execution Time
func measureTime(array []int, target int, sortedArray bool) float64 {
	var totalSortTime, totalSearchTime float64
	repeats := 10

	for i := 0; i < repeats; i++ {
		startSort := time.Now()
		if !sortedArray {
			sort.Ints(array)
		}
		totalSortTime += time.Since(startSort).Seconds()

		startSearch := time.Now()
		binarySearch(array, target)
		totalSearchTime += time.Since(startSearch).Seconds()
	}

	if sortedArray {
		return totalSearchTime / float64(repeats)
	}
	return (totalSortTime + totalSearchTime) / float64(repeats)
}

// Main Program
func main() {
	sizes := []int{100, 500, 1000, 5000, 10000}
	sortedTimes := make([]float64, len(sizes))
	unsortedTimes := make([]float64, len(sizes))

	for i, size := range sizes {
		array := generateArray(size, false)
		target := array[rand.Intn(len(array))]

		// Measure for sorted array
		sortedArray := generateArray(size, true)
		sortedTimes[i] = measureTime(sortedArray, target, true)

		// Measure for unsorted array
		unsortedTimes[i] = measureTime(array, target, false)
	}

	// Display Results

	fmt.Printf("%-10s %-20s %-20s\n", "Size", "Sorted Time (s)", "Unsorted Time (s)")
	for i, size := range sizes {
		fmt.Printf("%-10d %-20.10f %-20.10f\n", size, sortedTimes[i], unsortedTimes[i])
	}
}
