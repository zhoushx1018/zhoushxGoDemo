// Go supports _methods_ defined on struct types.

package main

import "fmt"

type rect struct {
	width, height int
}

// This `area` method has a _receiver type_ of `*rect`.
// 面积
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.

// 边长总和
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// 比较以下  modify1() 和 modify2() 函数
// 	modify2() 使用了 rect* 指针，因此才成功的修改了对象 width 值
func (r rect) modify1() {
	r.width = 8
}

func (r *rect) modify2() {
	r.width = 8
}

func main() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	// 比较以下  modify1() 和 modify2() 函数
	// 	modify2() 使用了 rect* 指针，因此才成功的修改了对象 width 值
	fmt.Println("", rp.width, rp.height)
	rp.modify1()
	fmt.Println("", rp.width, rp.height)
	rp.modify2()
	fmt.Println("", rp.width, rp.height)

}
