// 除了f mt 和 os ,我们还需要用到 bufio 来实现带缓冲输入（input）和输出(output)
// 读取用户的输入数据
// 我们怎样读取用户从键盘（控制台）输入的数据？输入指从键盘或其它标准输入（os.Stdin）读取数据。最简单的方法是使用fmt包里的Scan-或Sscan-系列函数，下面用个例子说明一下：
// // read input from the console:

package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	fmt.Println("Please input your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf(“%s %s”, &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}

// Scanln 将从标准输入的带有空格的字符串值保存到相应的变量里去，并以一个新行结束输入， Scanf做相同的工作，但它使用第一个参数指时输入格式， Sscan系列函数也是读取输入，但它是用来从字符串变量里读取，而不是从标准（os.Stdin）里读取
// 另外，我们也可以使用 bufio包里带缓冲的reader,例如
// //////////
// package main
// import (
//     "bufio"
//     "fmt"
//     "os"
// )
// var inputReader *bufio.Reader
// var input string
// var err error
// func main() {
//     inputReader = bufio.NewReader(os.Stdin)
//     fmt.Println("Please enter some input: ")
//     input, err = inputReader.ReadString('S') //func (b *Reader) ReadString(delim byte) (line string, err error) ,‘S’ 这个例子里使用S表示结束符，也可以用其它，如'\n'
//     if err == nil {
//         fmt.Printf("The input was: %s\n", input)
//     }
// }
// Please enter some input:
// abcd

// abc
// S
// The input was: abcd

// abc
// S

// 上例中，inputReader 是个指针，它指向一个 bufio类的Reader,然后在main函数里，通过了bufio.NewReader(os.Stdin)创建了一个buffer reader, 并联接到inputReader这个变量。bufio.NewReader（） 构造器的原型是这样的
// func NewReader(rd io.Reader) *Reader

// 任何符合io.Reader接口的对象（即实现了Read()方法对象）都可以作为bufio.NewReader()里的参数，并返回一个新的带缓冲的io.Reader, os.Stdin 符合这个条件。这个带缓冲的reader有一个方法ReadString(delim byte), 这个方法会一直读数据，直到遇到了指定的终止符，终止符将成为输入的一部分，一起放到buffer里去。
// ReadString 返回的是读到的字符串及nil;当读到文件的末端时，将返回把读到的字符串及io.EOF,如果在读到结束时没有发现所指定的结束符（delim byte）,将返回一个 err !=nil 。在上面的例子中，我们从键盘输入直到键入“S”。屏幕是标准输出os.Stdout，错误信息被写到os.Stderr,大多情况下，os.Stderr等同os.Stdout。
// 一般情况下，在GO的代码里，里省略了变量声明，而直接使用”:=“也声明，如：
// inputReader := bufio.NewReader(os.Stdin)
// input ,err :=inputReader.ReadString('\n')
// 下面的例子是使用了带关键字switch的，注意Go 的switch的几种形式以及unix和windows下不同的定界符。
// package main
// import (
//     "bufio"
//     "fmt"
//     "os"
// )
// func main() {
//     inputReader := bufio.NewReader(os.Stdin)
//     fmt.Println("please input your name:")
//     input, err := inputReader.ReadString('\n')
//     if err != nil {
//         fmt.Println("There ware errors reading,exiting program.")
//         return
//     }
//     fmt.Printf("Your name is %s", input)
//     //对unix:使用“\n”作为定界符，而window使用"\r\n"为定界符
//     //Version1
//     /*
//         switch input {
//         case "Philip\r\n":
//             fmt.Println("Welcome Philip!")
//         case "Chris\r\n":
//             fmt.Println("Welcome Chris!")
//         case "Ivo\r\n":
//             fmt.Println("Welcome Ivo!")
//         default:
//             fmt.Println("You are not welcome here! Goodbye!")
//         }
//     */
//     //version2
//     /*
//         switch input {
//         case "Philip\r\n":
//             fallthrough
//         case "Ivo\r\n":
//             fallthrough
//         case "Chris\r\n":
//             fmt.Printf("Welcome %s\n", input)
//         default:
//             fmt.Printf("You are not welcome here! Goodbye!")
//         }
//     */
//     //version3
//     switch input {
//     case "Philip\r\n", "Ivo\r\n", "Chris\r\n":
//         fmt.Printf("Welcome %s\n", input)
//     default:
//         fmt.Println("You are not welcome here! Goodbye!")
//     }
// }
