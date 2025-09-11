package main

import "fmt"

// 常量可以在整个包内被访问
const constant = 1

// 变量，作用域是包/函数作用域
var value string = "123"
var printCount int

func Basic() {
	a, b, c := 1, "b", "c"
	// = 是赋值语句
	a = 'a'
	b, c = c, b
	// := 是声明语句 , 声明则左侧必须要有新变量

	// no new variables on left side of :=compilerNoNewVar
	// b, c := returnChar()

	e, d := returnChar()
	b, c = returnString()
	Print(a, b, c, d, e)

	g := 1
	// p 是指向 int 类型的一个指针，他是一个内存地址
	// 从操作符来看， *p 相当于找 p 指针内存地址对应的变量， &g 是相反的操作， 返回 g 变量的内存地址
	// 从类型上看， *int 表示 指向 int 类型的指针
	// new 创建匿名函数， 返回的是地址,new 只是内置的函数，并不是一个保留字
	f := new(int)
	*f = 1
	p := &g
	// 对比值，对比内存地址
	Print(*p == g, p == &g, *f == *p, f == p) // T, T, T ,F
	h := escapeString()
	Print(h, *h) // 0x14000116ed0 , abc

}

func Print(values ...interface{}) {
	for _, value := range values {
		fmt.Printf("%d. ", printCount)
		fmt.Print(value)
		fmt.Print("\n")

		printCount++
	}
}

// 字符是 ‘’ ，字符串是 “”
func returnChar() (rune, rune) {
	return 'a', 'b'
}

func returnString() (string, string) {
	return "A", "B"
}

// 对于函数的入参
// 类型							传递方式
// 基本类型（int， string, bool ) 传递值的拷贝
// 指针类型（*)					传递指针的拷贝
// 切片（[])					传递切片头拷贝（包含底层数组的指针，长度，容量）
// map							传递 map 头拷贝
// 结构体						传递整个结构体的拷贝
// 接口 （interface{})			传递接口值的拷贝

func escapeString() *string {
	str := "abc"
	// 如果直接返回变量，将返回一个拷贝， 局部变量的内存占用将被 GC 销毁
	// 返回变量的指针， 则变量会被分配一个堆内存而不是栈内存， 他不会由作用域的结束和销毁，这叫逃逸
	// 最终由 GC 的可达性分析判断是否销毁
	return &str
}
