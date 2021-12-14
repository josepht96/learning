package main

import "fmt"

//InsertionSort uh
func insertionSort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		key := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > key {
				arr[j+1] = arr[j]
				arr[j] = key
			}
		}
	}
	for _, v := range arr {
		fmt.Printf("%d", v)
	}
}
func selectionSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		temp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}
	for _, v := range arr {
		fmt.Printf("%d", v)
	}
}
func main() {
	arr := []int{9, 4, 3, 5, 2, 1, 7}
	selectionSort(arr)
}
