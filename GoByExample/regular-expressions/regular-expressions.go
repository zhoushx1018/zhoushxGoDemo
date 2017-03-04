// Go offers built-in support for [regular expressions](http://en.wikipedia.org/wiki/Regular_expression).
// Here are some examples of  common regexp-related tasks
// in Go.

package main

import "bytes"
import "fmt"
import (
	"regexp"
	"runtime"
	"strings"
)
import "github.com/zhoushx1018/zhoushxGoDemo/GoByExample/print-file-line/fileline"

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

func main() {

	fmt.Println(FileLine(runtime.Caller(0)), "app init")

	// This tests whether a pattern matches a string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(FileLine(runtime.Caller(0)), "match=", match)

	// Above we used a string pattern directly, but for
	// other regexp tasks you'll need to `Compile` an
	// optimized `Regexp` struct.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Many methods are available on these structs. Here's
	// a match test like we saw earlier.
	fmt.Println(FileLine(runtime.Caller(0)), "r.MatchString=", r.MatchString("peach"))

	// This finds the match for the regexp.
	fmt.Println(FileLine(runtime.Caller(0)), "r.FindString=", r.FindString("peach punch"))

	// This also finds the first match but returns the
	// start and end indexes for the match instead of the
	// matching text.
	fmt.Println(FileLine(runtime.Caller(0)), "r.FindStringIndex=", r.FindStringIndex("peach punch"))

	// The `Submatch` variants include information about
	// both the whole-pattern matches and the submatches
	// within those matches. For example this will return
	// information for both `p([a-z]+)ch` and `([a-z]+)`.
	fmt.Println(FileLine(runtime.Caller(0)), "r.FindStringSubmatch=", r.FindStringSubmatch("peach punch"))

	// Similarly this will return information about the
	// indexes of matches and submatches.
	fmt.Println(FileLine(runtime.Caller(0)), "r.FindStringSubmatchIndex=", r.FindStringSubmatchIndex("peach punch"))

	// The `All` variants of these functions apply to all
	// matches in the input, not just the first. For
	// example to find all matches for a regexp.
	fmt.Println(FileLine(runtime.Caller(0)), "r.FindAllString=", r.FindAllString("peach punch pinch", -1))

	// These `All` variants are available for the other
	// functions we saw above as well.
	fmt.Println(FileLine(runtime.Caller(0)), "r.FindAllStringSubmatchIndex=", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Providing a non-negative integer as the second
	// argument to these functions will limit the number
	// of matches.
	fmt.Println(FileLine(runtime.Caller(0)), "r.FindAllString=", r.FindAllString("peach punch pinch", 2))

	// Our examples above had string arguments and used
	// names like `MatchString`. We can also provide
	// `[]byte` arguments and drop `String` from the
	// function name.
	fmt.Println(FileLine(runtime.Caller(0)), "r.Match=", r.Match([]byte("peach")))

	// When creating constants with regular expressions
	// you can use the `MustCompile` variation of
	// `Compile`. A plain `Compile` won't work for
	// constants because it has 2 return values.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(FileLine(runtime.Caller(0)), "MustCompile r=", r)

	// The `regexp` package can also be used to replace
	// subsets of strings with other values.
	fmt.Println(FileLine(runtime.Caller(0)), "r.ReplaceAllString=", r.ReplaceAllString("a peach", "<fruit>"))

	// The `Func` variant allows you to transform matched
	// text with a given function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(FileLine(runtime.Caller(0)), "r.ReplaceAllFunc=", string(out))
}
