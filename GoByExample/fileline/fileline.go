// 实现类似 c语言 的 __FILE__, __LINE__ 打印

package fileline

import "fmt"
import (
	"strings"
	"runtime"
)

//文件名、行号的获取和整理
// 实现类似 c 的 __FILE__, __LINE__
// 例子：
//	import . "zhoushxGoDemo/GoByExample/fileline"
// 	fmt.Println(FileLine(), "hello world")
func FileLine() (strRet string) {
	_, file, line, ok :=runtime.Caller(1)

	var s string
	if ok {
		slcTmp := strings.Split(file, "/")
		s = fmt.Sprintf("[%s %d] ", slcTmp[len(slcTmp)-1], line)
	}

	return s
}
