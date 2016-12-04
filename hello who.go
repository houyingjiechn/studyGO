package main

import (
	"fmt" //printfln函数所在包
	"os"  //切片os.Args所在包
	"strings" //字符串处理函数包
)

func main() { //主函数，无参数，默认无返回值，提取命令行输入程序名后的参数，gol
	who := "World!"       //快速变量声明
	if len(os.Args) > 1 { //如果切片os.Args长度大于1，
		who = strings.Join(os.Args[1:] "") //则把切片os.args第二个及后面的元素赋值字符串who中（提取程序参数）,
	}
	fmt.Println("Hello", who)
	fmt.Println(os.Args)
}
