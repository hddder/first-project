package main

import (
	"fmt"
)

//go语言中推荐使用驼峰式命名，如studentName

//单独声明变量
//var name string
//var ..

// 批量声明
var (
	name string
	age  int
	isok bool
)

func main() {
	name = "Hddder"
	age = 20
	isok = true
	//go语言中变量声明必须使用，否则无法编译过去

	fmt.Print(isok)             //在终端中输出内容
	fmt.Printf("name:%s", name) //格式化输出%:占位符，使用name变量去替换占位符
	fmt.Println(age)            //打印完指定内容后自动换行
	//go语言中不在乎缩进   ctrl+shift+P 搜format document自动格式化

}
