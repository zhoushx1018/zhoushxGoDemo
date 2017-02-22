// _Interfaces_ are named collections of method
// signatures.

// 接口，面向对象

package main

import "fmt"
import "math"

// Here's a basic interface for geometric shapes.
// "几何学" 接口
type geometry interface {
	area() float64			//面积
	perim() float64			//边长
}

// For our example we'll implement this interface on
// `rect` and `circle` types.
// 长方形
type rect struct {
	width, height float64
}

// 圆
type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.

// 长方形 rect 实现了 area() 和 perim() 方法， 因此 长方形 rect 是 “几何学”  geometry  这个接口的一个实现
//	使用这种方式，实现了 “几何学”  geometry  这个接口 的多态
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for `circle`s.
// 圆 circle 实现了 area() 和 perim() 方法， 因此 圆 circle 是 “几何学”  geometry  这个接口的一个实现
//	使用这种方式，实现了 “几何学”  geometry  这个接口 的多态
func (c circle) area() float64 {
	// 园的面积  =  π *  半径^2
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	// 圆的周长 =   2 * π  *  半径
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.

// measures 方法， 调用了 geometry 接口的 各种 方法(method)
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// The `circle` and `rect` struct types both
	// implement the `geometry` interface so we can use
	// instances of
	// these structs as arguments to `measure`.

	//分别传入  长方形 rect 和  圆 circle 对象， 实现多态
	measure(r)
	measure(c)
}