package main

import (
	"fmt"
)

//还有一种转换方式是switch测试。既然称之为switch测试，也就是说这种转换方式只能出现在switch语句中。
// 可以很轻松的将刚才用Comma-ok断言的例子换成由switch测试来实现：

type Html []interface{}

func main() {
	html := make(Html, 5)
	html[0] = "div"
	html[1] = "span"
	html[2] = []byte("script")
	html[3] = "style"
	html[4] = "head"
	for index, element := range html {
		switch value := element.(type) {
		case string:
			fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
		case []byte:
			fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
		case int:
			fmt.Printf("invalid type\n")
		default:
			fmt.Printf("unknown type\n")
		}
	}
}