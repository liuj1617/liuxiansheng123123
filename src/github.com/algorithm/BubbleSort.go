package main

import "fmt"

//冒泡
func bubbling(array []int) {
	length := len(array)
	for i := 0; i < length-1; i++ {
		if array[i] > array[i+1] {
			tempInt := array[i]
			array[i] = array[i+1]
			array[i+1] = tempInt
		}
	}
}

//冒泡排序
func bubbleSort(array []int) {
	if array == nil || len(array) == 0 {
		fmt.Println("array is nil")
		return
	}
	length := len(array)
	for i := length; i >= 1; i-- {
		bubbling(array)
	}
}

//冒泡排序
func main() {
	bubblingArray := []int{9, 0, 8, 6, 4, 2, 3, 1, 5, 7} //需要排序的数组
	fmt.Println(bubblingArray)                           //打印排序前数组
	bubbleSort(bubblingArray)                            //排序
	fmt.Println(bubblingArray)                           //打印排序后数组
}
