package main

import "fmt"

//寻找数组中最大元素的下标并返回
func findMaxIndex(array []int, n int) int {
	max := array[0] //最大值默认第一位
	index := 0
	for i := 0; i < n; i++ {
		if array[i] > max { //如果有比第一位大的数字，则最大值赋值成最大值
			max = array[i]
			index = i //记录最大值下标并返回
		}
	}
	return index
}

//选择排序
func selectionSort(array []int) {
	if array == nil || len(array) == 0 { //判空
		fmt.Println("array is nil")
		return
	}
	length := len(array) //数组长度
	//找到最大的下标，将最大的数字与数组最后一位数字进行交换
	for i := length; i > 1; i-- {
		index := findMaxIndex(array, i) //寻找数组前i位最大数字下标
		tempInt := array[i-1]           //最后一位存储给临时变量
		array[i-1] = array[index]       //本次查找数组最后一位赋值为最大值
		array[index] = tempInt          //原最大值位置赋值成最后一位
	}
}

//选择排序
func main() {
	selectionArray := []int{0, 8, 6, 4, 2, 3, 1, 9, 5, 7} //需要排序的数组
	fmt.Println(selectionArray)                           //打印排序前数组
	selectionSort(selectionArray)                         //排序
	fmt.Println(selectionArray)                           //打印排序后数组
}
