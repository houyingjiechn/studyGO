package main

import "fmt"

func adder() func(int) int { //函数adder无参数，返回闭包函数func(int)，返回值类型int
	sum := 0
	return func(x int) int { //函数返回闭包函数的值，闭包函数接收参数x，返回值int
		sum += x
		return sum //闭包函数返回sum
	}
}

func main() {
	pos, neg := adder(), adder() //pos,neg的类型为函数的值
	for i := 0; i < 10000; i++ {
		fmt.Println(
			pos(i), //使用参数调用pos，函数封装
			neg(-2*i),
		)
	}
}
