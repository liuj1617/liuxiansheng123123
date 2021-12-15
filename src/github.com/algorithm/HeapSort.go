package main

import "fmt"

func swap(array []int, i int, j int) {
	tempInt := array[i]
	array[i] = array[j]
	array[j] = tempInt
}

func heapify(array []int, n int, i int) {
	if i >= n {
		return
	}
	c1 := 2*i + 1
	c2 := 2*i + 2
	max := i
	if c1 < n && array[c1] > array[max] {
		max = c1
	}
	if c2 < n && array[c2] > array[max] {
		max = c2
	}
	if max != i {
		swap(array, max, i)
		heapify(array, n, max)
	}
}

func build_heap(array []int, n int) {
	last_node := n - 1
	parent := (last_node - 1) / 2
	for i := parent; i >= 0; i-- {
		heapify(array, n, i)
	}
}

func heapSort(array []int, n int) {
	build_heap(array, n)
	for i := n - 1; i >= 0; i-- {
		swap(array, i, 0)
		heapify(array, i, 0)
	}
}

func main() {
	heapArray := []int{0, 8, 6, 4, 2, 3, 1, 9, 5, 7} //需要排序的数组
	fmt.Println(heapArray)                           //打印排序前数组
	heapSort(heapArray, len(heapArray))              //排序
	fmt.Println(heapArray)                           //打印排序后数组
}
