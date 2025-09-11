package packageInit

import "fmt"

// init 函数在被 import 的时候调用
// 流程： 开始  - 初始化全局变量 - import 包触发 init - 触发当前包 init - 触发 main 函数
func init() {
	fmt.Println("这是 packageInit 的 init 函数 1")
}

// 如果有多个 init 按顺序执行
func init() {
	fmt.Println("这是 packageInit 的 init 函数 2")
}

func PackageInit() {
	fmt.Println("这是 packageInit 的函数逻辑")
}
