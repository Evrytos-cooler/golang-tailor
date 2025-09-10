// 包声明：main 包表示这是一个可执行程序（而非库）
package main

// 导入依赖包：fmt 包用于输入输出（如打印内容到终端）
import (
	"fmt"
	"os"
	"strings"
)

// 自定义的 Hello 函数
func Hello() string {
	return "Hello from my custom function!"
}

var abc = "123"

// main 函数：程序的入口函数（必须有，且无参数无返回值）
func main() {
	fmt.Println("123")
	fmt.Println(strings.Join(os.Args, " "))
}
