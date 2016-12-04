package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := exec.Command("cmd", "/c", "pause", "D:\\a.txt")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
