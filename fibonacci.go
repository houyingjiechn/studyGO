package main

import (
	"fmt"
	"math"
)

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func(int) int {
	sq5 := math.Sqrt(5)
	padd2 := (1 + sq5) / 2
	pdim2 := (1 - sq5) / 2
	resq5 := 1 / sq5
	return func(n int) int {
		fn := float64(n)
		an := resq5 * (math.Pow(padd2, fn) - math.Pow(pdim2, fn))
		return int(an)
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Println(f(i))
	}
}
