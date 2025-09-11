package main

import (
	"fmt"
	"go-feature-learning/packageInit"
)

var global = "global"

func main() {
	fmt.Printf("这是 main 包的 main 函数逻辑 %s \n", global)
	packageInit.PackageInit()
	// FmtLearning()
	// ParallelFetch([]string{"https://baidu.com", "https://google.com", "https://github.com"})
	// WebFetch()
	// Print(constant,value)
	// Basic()
	// TypeFunc()
}
