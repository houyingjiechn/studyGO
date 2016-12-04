package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday() //按周计时，今天周几
	sat := time.Saturday
	defer fmt.Println(today, sat, "defer") //defer延迟执行，直到上层函数结束才执行
	switch sat {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

}
