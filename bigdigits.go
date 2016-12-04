package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//定义二维切片，字库

var bigDigits = [][]string{
	{
		" 000 ",
		"0   0",
		"0   0",
		"0   0",
		" 000 ",
	},
	{
		"  1 ",
		" 11 ",
		"  1 ",
		"  1 ",
		"1111",
	},
	{
		" 222 ",
		"2   2",
		"   2 ",
		"  2  ",
		"22222",
	},
	{
		" 333 ",
		"3   3",
		"  33 ",
		"3   3",
		" 333 ",
	},
	{
		"   4 ",
		"  44 ",
		" 4 4 ",
		"44444",
		"   4 ",
	},
	{
		"55555",
		"5    ",
		"5555 ",
		"    5",
		"5555 ",
	},
	{
		" 666 ",
		"6    ",
		"6666 ",
		"6   6",
		" 666 ",
	},
	{
		"77777",
		"   7 ",
		"  7  ",
		" 7   ",
		"7    ",
	},
	{
		" 888 ",
		"8   8",
		" 888 ",
		"8   8",
		" 888 ",
	},
	{
		" 999 ",
		"9   9",
		" 999 ",
		"  9  ",
		" 9   ",
	},
}

//主函数

func main() {

	//判断程序名后有没有参数
	if len(os.Args) == 1 { // os.Args数组存储了文件名信息，os.Args[0]存储了文件名
		fmt.Printf("usage:%s <whole-number>\n", filepath.Base(os.Args[0])) //%s 指代 参数，filepath.Base(os.Args[0]) 返回文件路径基础名（文件名）
		os.Exit(1)                                                         //程序向操作系统返回1 ，即程序失败
	}

	sod := os.Args[1] //提取参数

	for row := range bigDigits[0] { //row=行数，用range遍历切片 （for循环5次）
		line := "" //初始化列

		for column := range sod { //提取参数数字，循环次数等于参数数字位数
			digit := sod[column] - '0' //UTF-8码-0(48)，单引号表达字符，字符可以和所有整数类型上下文兼容
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " " //逐位提取数字
			} else {
				log.Fatal("invaild whole number") //错误日志
			}
		}

		fmt.Println(line)
	}

}
