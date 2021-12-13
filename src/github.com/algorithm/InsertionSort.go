package main

import "fmt"

//插入排序，类似于打扑克
func insert(array []int, n int) {
	//key := array[n]
	//i := 0
	//for {
	//	if
	//	array[n] = key
	//		n--
	//	i++
	//	if !(array[n-1] < key) {
	//		break
	//	}
	//}

}

func insertSort(array []int, n int) {
	for i := 1; i < n; i++ {

	}

}

//插入排序
func main() {
	selectionArray := []int{0, 8, 6, 4, 2, 3, 1, 9, 5, 7} //需要排序的数组
	fmt.Println(selectionArray)                           //打印排序前数组
	selectionSort(selectionArray)                         //排序
	fmt.Println(selectionArray)                           //打印排序后数组
}
