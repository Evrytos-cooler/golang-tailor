package main

import "fmt"

// type 其实就是给类型起别名，不同类型之间不能比较
type A int
type B int

// 定义类型的方法
//
//	是 A 类型的方案， a 相当于 this ， toString() 函数名和参数， string 返回值
func (this A) toString() string {
	return fmt.Sprintf("this is type A %d", this)
}
func (this B) toString() string {
	return fmt.Sprintf("this is type B %d", this)
}

func TypeFunc() {
	var a A = 1
	var b B = 2
	fmt.Println(a.toString())
	fmt.Println(b.toString())
	// invalid operation: a == b (mismatched types A and B)compilerMismatchedTypes
}
