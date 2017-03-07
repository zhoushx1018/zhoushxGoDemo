package main

import (
	"fmt"
)

//其实Comma-ok断言还支持另一种简化使用的方式：value := element.(T)。
// 但这种方式不建议使用，因为一旦element.(T)断言失败，则会产生运行时错误。如：

func main() {
	var val interface{} = "good"
	fmt.Println(val.(string))
	// fmt.Println(val.(int))
}