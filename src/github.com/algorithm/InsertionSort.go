package main

import "fmt"

//插入排序，类似于打扑克
func insert(array []int, n int) {
	key := array[n]
	i := n
	for array[i-1] > key {
		array[i] = array[i-1]
		i--
		if i == 0 {
			break
		}
	}
	array[i] = key
}

//插入排序
func insertionSort(array []int, n int) {
	for i := 1; i < n; i++ {
		insert(array, i)
	}
}

//插入排序
func main() {
	insertionArray := []int{0, 8, 6, 4, 2, 3, 1, 9, 5, 7} //需要排序的数组
	fmt.Println(insertionArray)                           //打印排序前数组
	insertionSort(insertionArray, len(insertionArray))    //排序
	fmt.Println(insertionArray)                           //打印排序后数组
}
