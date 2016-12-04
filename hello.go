package main

import (
	"fmt" //包含 fmt 包，使用Println函数
)

func main() {
	var ( //定义变量,不可定义未使用的变量
		i, k float64 = 10, 20
	)
	//	i := 10
	//	k := 20

	const ( //定义常量
		a  = 1 << iota //a=1b iota=0 编译器移位计算赋值，每次引用预定义常量iota=0
		b              //b=1b iota=1 枚举
		c              //c=100b iota=2 c=4
		d  = iota      //iota=3
		e              //iota=4
		f              //iota=5 f=5
		Pi = 3.14      //以大写字符开头的常量可以被其他包访问，小写字符开头的常量为包内私有
	)
	i, k = k, i //变量交换

	l := i*Pi + k + e // := 自动判断变量类型赋值
	fmt.Println("hello world!", l, c, f)
}
