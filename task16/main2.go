package main

import "fmt"

func main() {
	arr := [5]int{1, 0, 5, 11, 9}
	
	fmt.Println(QuickSort(arr[:], 0, len(arr)-1))
}

func QuickSort(arr []int, left int, right int) []int {
	if left < right {
		p, arr := partition(arr, left, right)
		arr = QuickSort(arr, left, p-1)
		arr = QuickSort(arr, p+1, right)
	}
	return arr 
}

func partition(arr []int, left int, right int) (int, []int) {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[right], arr[i] = arr[i], arr[right]
	return i, arr
}