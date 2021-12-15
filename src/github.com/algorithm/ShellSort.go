package main

import "fmt"

func shellSort(array []int) {
	n := len(array)
	var i, j, gap, tempInt int
	for gap = n / 2; gap > 0; gap = gap / 2 {
		for i = gap; i < n; i++ {
			tempInt = array[i]
			for j = i; j >= gap && tempInt < array[j-gap]; j -= gap {
				array[j] = array[j-gap]
			}
			array[j] = tempInt
		}
	}
}

//希尔排序
func main() {
	shellArray := []int{0, 8, 6, 4, 2, 3, 1, 9, 5, 7} //需要排序的数组
	fmt.Println(shellArray)                           //打印排序前数组
	shellSort(shellArray)                             //排序
	fmt.Println(shellArray)                           //打印排序后数组
}
