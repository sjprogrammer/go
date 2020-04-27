/**
 * 指针
 * 指针分等级:
 * 			一般变量默认是零级指针,根据*的数量来区别指针的等级,根据指针的等级来接受指针变量;
 *			比如一级指针接受零级指针,二级指针接受一级指针
 *
 *
 */
package main

import "fmt"

func init() {
	fmt.Println("指针语法")
}

func pointDemo() {
	// pointAddress()
	pointType()
}

// &获取变量的存储地址
func pointAddress() {
	num := 1
	fmt.Printf("取址符&:%v \n", &num)
}

// 申明指针(基本类型变量)
func pointType() {
	var num *int    // 1级指针
	var num1 **int  // 2级指针
	var num2 ***int // 3级指针
	val := 1
	num = &val
	num1 = &num
	num2 = &num1
	fmt.Printf("取址符&:%v \n", *num)
	fmt.Printf("取址符&:%v \n", *num1)
	fmt.Printf("取址符&:%v \n", *num2)
	val = 10
	fmt.Printf("取址符&:%v \n", *num)
}

// 声明指针(数组变量)
func pointArray() {
	var arr [2]*int
	fmt.Printf("取址符&:%v", &arr)
}
