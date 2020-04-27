package main

import "fmt"

// 数组操作
func init() {
	arr := [1]int{1}
	fmt.Printf("打印数组:%v", arr)
}

/**
 * 声明
 */
func varOne() {
	// 第一种 固定数组长度,并赋值默认值
	var arr = [2]int{2, 1}
	fmt.Println(arr)
	// 第二种 固定数组长度,没有默认值
	var arrTwo [2]int
	fmt.Println(arrTwo)
}

/**
 * 数据第一种循环 for index; length;rule {}
 */
func loopOne() {
	arr := [2]int{1, 9}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%v]=%v \n", i, arr[i])
	}
}

/**
 * 数据第二种循环 range
 */
func loopTwo() {
	arr := [2]int{4, 1}
	for index, value := range arr {
		fmt.Printf("index=%v;value=%v", index, value)
	}
}
