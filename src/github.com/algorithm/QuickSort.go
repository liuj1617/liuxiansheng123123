package main

import "fmt"

func Partition(array []int, first int, end int) int {
	i := first
	j := end
	var tempInt int
	for i < j {
		for i < j && array[i] < array[j] {
			j--
		}
		if i < j {
			tempInt = array[i]
			array[i] = array[j]
			array[j] = tempInt
			i++
		}
		for i < j && array[i] < array[j] {
			i++
		}
		if i < j {
			tempInt = array[i]
			array[i] = array[j]
			array[j] = tempInt
			j--
		}
	}
	return i
}

func QuickSort(array []int, first int, end int) {
	var pivot int
	if first < end {
		pivot = Partition(array, first, end)
		QuickSort(array, first, pivot-1)
		QuickSort(array, pivot+1, end)
	}
}

//快速排序
func main() {
	QuickArray := []int{0, 8, 6, 4, 2, 3, 1, 9, 5, 7} //需要排序的数组
	fmt.Println(QuickArray)                           //打印排序前数组
	QuickSort(QuickArray, 0, len(QuickArray)-1)       //快速排序
	fmt.Println(QuickArray)                           //打印排序后数组
}
