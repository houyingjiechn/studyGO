package main

import (
	"fmt"
)

func main() {

	var str string

	// 字符串

	str = "hello world!"
	ch := str[0]

	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)

	fmt.Printf("\n")

	n := len(str)

	//遍历字符串

	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch)
	}

	fmt.Printf("\n")

	for i, ch := range str {
		fmt.Println(i, ch) //ch的类型为rune
	}

	fmt.Printf("\n")

	// 数组

	var array [3][3][3][3]int //定义四维数组

	array[0][0][0][0] = 10 // 试图修改数组的第一个元素

	fmt.Println("array values:", array) //显示数组所有元素

	for i := 0; i < len(array); i++ { //遍历数组
		fmt.Println("Element", i, "of array is", array[i])
	}

	for i, v := range array { //遍历数组关键字range ， 有两个返回值，元素下标和元素的值
		fmt.Println("Array element[", i, "]=", v)
	}

	fmt.Printf("\n")

	// 根据数组定义数组切片
	var array2 [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//定义数组切片，数组[first : last],[:5]=[0:5]，[:]=所有元素，数组切片可大于数组长度
	var slice []int = array2[:5]

	/*直接定义数组切片，用make函数
	mySlice1 := make([]int, 5)       //整型数组切片，5个值为0的初始元素
	mySlice2 := make([]int, 5, 10)   //整型数组切片，5个值为0的初始元素，预留10个元素的存储空间
	mySlice3 := []int{1, 2, 3, 4, 5} //直接创建并初始化5个元素的数组切片
	*/

	slice = append(slice, 1, 2, 3) //使用append函数在切片尾端添加元素
	//slice = append(slice, slice2...)  使用append函数在切片尾端添加一个切片，尾端的...要保留
	//数组切片会自动处理存储空间不足的问题

	//以数组切片创建数组切片，新切片的长度可以大于原切片但要小于原切片的存储空间
	//oldSlice := []int{1, 2, 3, 4, 5}
	//newSlice := oldSlice[:3] // 基于oldSlice的前3个元素构建新数组切片

	/*复制切片 cpoy()
		slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	*/

	fmt.Println("Elements of myArray: ") //遍历数组

	for _, v := range array2 {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice: ")

	for _, v := range slice {
		fmt.Print(v, " ")
	}

	fmt.Printf("\n")

	for i, v := range slice {
		fmt.Println("\nmySlice[", i, "] =", v)
	}

	fmt.Println()

	fmt.Println("len(mySlice):", len(slice)) //len函数返回切片元素个数
	fmt.Println("cap(mySlice):", cap(slice)) //cap函数返回切片空间大小
}
