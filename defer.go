package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //defer延迟执行，但结果被压入一个栈中，按照顺序被调用
	}

	fmt.Println("done")
}
