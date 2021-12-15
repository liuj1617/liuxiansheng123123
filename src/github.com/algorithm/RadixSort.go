package main

import "fmt"

func radixSort(array []int) {
	if len(array) < 2 {
		return
	}
	max := array[0]
	dataLen := len(array)
	for i := 1; i < dataLen; i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	maxDigit := 0
	for max > 0 {
		max = max / 10
		maxDigit++
	}
	divisor := 1
	bucket := [10][20]int{{0}}
	count := [10]int{0}
	var digit int
	for i := 1; i <= maxDigit; i++ {
		for j := 0; j < dataLen; j++ {
			tmp := array[j]
			digit = (tmp / divisor) % 10
			bucket[digit][count[digit]] = tmp
			count[digit]++
		}
		k := 0
		for b := 0; b < 10; b++ {
			if count[b] == 0 {
				continue
			}
			for c := 0; c < count[b]; c++ {
				array[k] = bucket[b][c]
				k++
			}
			count[b] = 0
		}
		divisor = divisor * 10
	}

}

//基数排序
func main() {
	radixArray := []int{0, 8, 6, 4, 2, 3, 1, 9, 5, 7} //需要排序的数组
	fmt.Println(radixArray)                           //打印排序前数组
	radixSort(radixArray)
	fmt.Println(radixArray) //打印排序后数组
}
