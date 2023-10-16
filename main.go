package main

import (
	"fmt"
	"sorting/internal/bubble"
	"sorting/internal/merge"
)

func main() {
	arr := []int{5, 2, 6, 3, 1, 4}
	fmt.Println("The given unsorted array is:", arr)

	mergeSorted := merge.MergeSort(arr)
	fmt.Println("The merge-sorted array is:", mergeSorted)

	fmt.Println("The given unsorted array is:", arr)
	bubble.BubbleSort(arr)
	fmt.Println("The bubble-sorted array is:", arr)
}
