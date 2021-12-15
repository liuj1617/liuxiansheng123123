package main

import "fmt"

func merge(array []int, left int, mid int, right int) {
	leftSize := mid - left
	rightSize := right - mid + 1
	leftArray := make([]int, leftSize)
	rightArray := make([]int, rightSize)
	var i, j, k int
	//左边数组赋值
	for i = left; i < mid; i++ {
		leftArray[i-left] = array[i]
	}
	//右边数组赋值
	for i = mid; i <= right; i++ {
		rightArray[i-mid] = array[i]
	}

	//归并到原始数组
	i = 0
	j = 0
	k = left
	for i < leftSize && j < rightSize {
		if leftArray[i] < rightArray[j] {
			array[k] = leftArray[i]
			i++
			k++
		} else {
			array[k] = rightArray[j]
			j++
			k++
		}
	}

	for i < leftSize {
		array[k] = leftArray[i]
		i++
		k++
	}

	for i < rightSize {
		array[k] = rightArray[j]
		j++
		k++
	}
}

func mergeSort(array []int, left int, right int) {
	if left == right {
		return
	} else {
		mid := (left + right) / 2
		mergeSort(array, left, mid)
		mergeSort(array, mid+1, right)
		merge(array, left, mid+1, right)
	}
}

//归并排序
func main() {
	array := []int{2, 8, 9, 10, 4, 5, 6, 7, 1, 0, 3}
	mergeSort(array, 0, len(array)-1)
	fmt.Println(array)

}
