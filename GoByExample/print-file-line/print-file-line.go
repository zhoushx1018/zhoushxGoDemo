// 实现类似 c语言 的 __FILE__, __LINE__ 打印

package fileline

import "fmt"
import (
	//"runtime"
	"strings"
)

//文件名、行号的获取和整理
// 实现类似 c 的 __FILE__, __LINE__
// 例子：
// 	fmt.Println(FileLine(runtime.Caller(0)), "hello world")
func FileLine(pc uintptr, file string, line int, ok bool) (strRet string) {
	var s string
	if ok {
		slcTmp := strings.Split(file, "/")
		s = fmt.Sprintf("[%s %d] ", slcTmp[len(slcTmp)-1], line)
	}

	return s

}
//
//func main() {
//
//	fmt.Println(FileLine(runtime.Caller(0)), "hello world")
//
//}
