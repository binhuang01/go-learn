package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	/*
		range会产生一对值，索引和索引处对应的值，range语法规定，要处理元素，
		必须处理索引。把索引赋值给一个临时变量，然后忽略它的值，go不允许使用无用
		的局部变量。解决办法是用空标识符，即_。
	*/
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
