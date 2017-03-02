// Sometimes we'll want to sort a collection by something
// other than its natural order. For example, suppose we
// wanted to sort strings by their length instead of
// alphabetically. Here's an example of custom sorts
// in Go.

package main

import "sort"
import "fmt"

// In order to sort by a custom function in Go, we need a
// corresponding type. Here we've created a `ByLength`
// type that is just an alias for the builtin `[]string`
// type.

//定义一个新的类型 "ByLength",  一个 字符串的 slice
type ByLength []string

// We implement `sort.Interface` - `Len`, `Less`, and
// `Swap` - on our type so we can use the `sort` package's
// generic `Sort` function. `Len` and `Swap`
// will usually be similar across types and `Less` will
// hold the actual custom sorting logic. In our case we
// want to sort in order of increasing string length, so
// we use `len(s[i])` and `len(s[j])` here.

//(类型 "ByLength") 的函数定义
//  用以实现  sort.Interface 接口；  相关函数包括  Len, Less, Swap
//	返回 slice 的长度
func (s ByLength) Len() int {
	return len(s)
}
//(类型 "ByLength") 的函数定义
//	交换 slice 中，i和j 的值
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
//(类型 "ByLength") 的函数定义
//	根据 字符串的长度，比大小
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// With all of this in place, we can now implement our
// custom sort by casting the original `fruits` slice to
// `ByLength`, and then use `sort.Sort` on that typed
// slice.

//使用 sort.Sort(interface) 的多态调用
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}