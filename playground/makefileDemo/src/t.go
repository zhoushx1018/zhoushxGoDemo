package main

import "fmt"

func main() {
	//    fmt.Printf("Hello, world!\n")
	fmt.Printf("1111111Hello, world!\n")
}

// package main

//import "fmt"
//
//// -------------------------------------------------------------
//
//type IReadWriter interface {
//	Read(buf *byte, cb int) int
//	Write(buf *byte, cb int) int
//}
//
//// -------------------------------------------------------------
//
//type A struct {
//	a int
//}
//
//func NewA(params int) *A {
//	fmt.Println("NewA:", params)
//	return &A{params}
//}
//
//func (this *A) Read(buf *byte, cb int) int {
//	fmt.Println("A_Read:", this.a)
//	return cb
//}
//
//func (this *A) Write(buf *byte, cb int) int {
//	fmt.Println("A_Write:", this.a)
//	return cb
//}
//
//// -------------------------------------------------------------
//
//type B struct {
//	A
//}
//
//func NewB(params int) *B {
//	fmt.Println("NewB:", params)
//	return &B{A{params}}
//}
//
//func (this *B) Write(buf *byte, cb int) int {
//	fmt.Println("B_Write:", this.a)
//	return cb
//}
//
//func (this *B) Foo() {
//	fmt.Println("B_Foo:", this.a)
//}
//
//// -------------------------------------------------------------
//
//func main() {
//	var p IReadWriter = NewB(18)
//	iRet := p.Read(nil, 10)
//	fmt.Println("iRet1:", iRet)
//
//	iRet = p.Write(nil, 10)
//	fmt.Println("iRet2:", iRet)
//
//	p = NewA(98)
//	iRet = p.Read(nil, 90)
//	fmt.Println("iRet1:", iRet)
//
//	iRet = p.Write(nil, 90)
//	fmt.Println("iRet2:", iRet)
//
//}

//
//// 闭包
//package main
//
//import "fmt"
//
//// fib returns a function that returns
//// successive Fibonacci numbers.
//func fib() func() int {
//	//	a, b := 0, 1
//	a := 0
//	return func() int {
//		//		a, b = b, a+b
//		a = a + 1
//		return a
//	}
//}
//
//func main() {
////	if 1 == 1 {
////		f := fib()
////		var i int64
////		for i = 1; i < 100; i++ {
////			println(f(), f(), f(), f(), f())
////		}
////
////	}
//
//	f1 := fib()
//	fmt.Println(f1())
//	fmt.Println(f1())
//	fmt.Println(f1())
//	fmt.Println(f1())
//
//	fmt.Println(f1(), f1(), f1())
//
//	fmt.Println(f1(), f1())
//
//	f2 := fib()
//	fmt.Println(f2(), f2())
//
//
//}
//

//
//package main
//
//import (
//	"flag"
//	"fmt"
//	"os"
//)
//
//
//
//const (
//	Space   = " "
//	NewLine = "\n"
//)
//
//var omitNewLine = flag.Bool("n", false, "dont't print final newline")
//
//
//func main() {
//	flag.Parse()
//	var s string = ""
//	fmt.Println("gogogogoggo")
//	for i := 0; i < flag.NArg(); i++ {
//		fmt.Println(i)
//
//		if i > 0 {
//			s += Space
//		}
//		s += flag.Arg(i)
//	}
//
//
//
//	if !*omitNewLine {
//		s += NewLine
//	}
//
//	os.Stdout.WriteString(s)
//
//}
//
//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//func sum(a []int) int { // returns an int
//	s := 0
//	for i := 0; i < len(a); i++ {
//		s += a[i]
//	}
//	return s
//}
//
//func main() {
//	s := "hello"
//	if s[1] != 'e' {
//		fmt.Println("not e .............")
//		os.Exit(1)
//	}
//	s = "good bye"
//
//	fmt.Println(s)
//
//	var p *string = &s
//	*p = "ciao"
//
////	var aTmp = [6]int{1,2,3,4,5}
////	iRet := sum([6]int{1,2,3,4,5,6}[:])
//	s1 := sum([]int{1,2,3})
//	fmt.Println(s1)
//
//}

//package main

//import "fmt"

//func main() {
//	i, o := make(chan string), make(chan string)
//	go func() {
//		for {
//			o <- <-i
//		}
//	}()
//	// i <- 0xBabeFac

//	// var str string = "abcd"
//	str := "abcd"
//	i <- str
//	fmt.Println(<-o)
//}

//
//package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//type Hello struct{}
//
//func (h Hello) ServeHTTP(
//	w http.ResponseWriter,
//	r *http.Request) {
//	fmt.Fprint(w, "Hello world! by  go web package")
//}
//
//func main() {
//	var h Hello
//	http.ListenAndServe("192.168.9.26:4000", h)
//}

//
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	fmt.Println("Happy pi:[", math.Pi, "] Day")
//}
//

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	fmt.Printf("Now you have %g problems. \n",
//		math.Nextafter(2, 3))
//
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	fmt.Println(math.Pi)
//}

//A tour of go-6 函数
//package main
//
//import "fmt"
//
//func add(x int, y int) int {
//	return x + y
//}
//
//func main() {
//	fmt.Println(add(42, 13))
//}

////A tour of go-8 函数
//// 函数有多个返回值
//package main
//
//import "fmt"
//
//func swap(x, y string) (string, string) {
//	return y, x
//}
//
//func main() {
//	a, b := swap("hello", "world")
//	fmt.Println(a, b)
//}

//////A tour of go-9 函数
//// 函数有多个返回值
//package main
//
//import "fmt"
//
//func split(sum int) (x, y int) {
//	x = sum * 4/9
//	y = sum - x
//	return
//}
//
//func main() {
//	fmt.Println(split(17))
//}

//
//package main
//
//import "fmt"
//
//var x, y, z int
//var c, python, java bool
//
//func main() {
//	fmt.Println(x, y, z, c, python, java)
//}

//package main
//
//import "fmt"
//
//const Pi = 3.14
//
//func main() {
//	const World = "世界"
//	fmt.Println("Hello", World)
//	fmt.Println("Happy", Pi, "Day")
//
//	const Truth = true
//	fmt.Println("Go rules?", Truth)
//}
//
//package main
//
//import "fmt"
//
//const (
//	Big   = 1 << 100
//	Small = Big >> 99
//)
//
//func needInt(x int) int { return x*10 + 1 }
//func needFloat(x float64) float64 {
//	return x * 0.1
//}
//
//func main() {
//	var fA float64 = Big
//	fmt.Println("fA [", fA, "]")
//
//	var fB float64 = Small
//	fmt.Println("fB [", fB, "]")
//
//	fmt.Println(needInt(Small))
//	fmt.Println(needFloat(Small))
//	fmt.Println(needFloat(Big))
//}

//package main
//
//import "fmt"
//
//func main() {
//	sum := 0
//	for i := 0; i < 10; i++ {
//		sum += i
//	}
//	fmt.Println(sum)
//}
//

//
//package main
//
//func main() {
//	for  {
//	}
//}
//
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func pow(x, n, lim float64) float64 {
//	if v := math.Pow(x, n); v < lim {
//		return v
//	}
//	return lim
//}
//
//func main() {
//	fmt.Println(
//		pow(3, 2, 10),
//		pow(3, 3, 20),
//	)
//}

//
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func pow(x, n, lim float64) float64 {
//	if v := math.Pow(x, n); v < lim {
//		return v
//	} else {
//		fmt.Printf("%g >= %g\n", v, lim)
//	}
//        // 不能在这里使用 v，因此
//	return lim
//}
//
//func main() {
//	fmt.Println(
//		pow(3, 2, 10),
//		pow(3, 3, 20),
//	)
//}
//
//
//package main
//
//import (
//	"math/cmplx"
//	"fmt"
//)
//
//var (
//	ToBe bool = false
//	MaxInt uint64 = 1<<64 - 1
//	z complex128 = cmplx.Sqrt(-5+12i)
//)
//
//func main() {
//	const f = "%T(%v)\n"
//	fmt.Printf(f, ToBe, ToBe)
//	fmt.Printf(f, MaxInt, MaxInt)
//	fmt.Printf(f, z, z)
//}
//
//package main
//
//import "fmt"
//
//type Vertex struct {
//	X int
//	Y int
//}
//
//func main() {
//	p := Vertex{1, 2}
//	q := &p
//	q.X = 1e9
//	p.Y = 3
//	fmt.Println(p)
//}

//
//package main
//
//import "fmt"
//
//type Vertex struct {
//	X, Y int
//}
//
//var (
//	p = Vertex{1, 2}  // has type Vertex
//	q = &Vertex{1, 2} // has type *Vertex
//	r = Vertex{X: 1}  // Y:0 is implicit
//	s = Vertex{}      // X:0 and Y:0
//)
//
//func main() {
//	fmt.Println(p, q, r, s)
//}
//
//package main
//
//import "fmt"
//
//type Vertex struct {
//	X, Y int
//}
//
//func main() {
//	v := new(Vertex)
//	fmt.Println(v)
//	v.X, v.Y = 11, 9
//	fmt.Println(v)
//}

//
//
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	hypot := func(x, y float64) float64 {
//		return math.Sqrt(x*x + y*y)
//	}
//
//	fmt.Println(hypot(3, 4))
//}
//
//
//package main
//
//import "fmt"
//
//func main() {
//	m := make(map[string]int)
//
//	m["Answer"] = 42
//	fmt.Println("The value:", m["Answer"])
//
//	m["Answer"] = 48
//	fmt.Println("The value:", m["Answer"])
//
//	v0, ok_0 := m["Answer"]
//	fmt.Println( "vo=", v0, "|ok_0=", ok_0 );
//
//	delete(m, "Answer")
//	fmt.Println("The value:", m["Answer"])
//
//	v, ok := m["Answer"]
//	fmt.Println("The value:", v, "Present?", ok)
//}

//package main
//
//import "fmt"
//
//func adder() func(int) int {
//	sum := 0
//	return func(x int) int {
//		sum += x
//		return sum
//	}
//}
//
//func main() {
//	pos, neg := adder(), adder()
//	for i := 0; i < 10; i++ {
//		fmt.Println(
//			pos(1),
//			neg(-2),
//		)
//	}
//}

//index  39  range

//package main
//
//import "fmt"
//
//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
//
//func main() {
//	for i, v := range pow {
//	    fmt.Printf("2**%d = %d\n", i, v)
//	}
//
//	fmt.Printf("__________________\n" )
//
//	for i1 := range pow {
//	    fmt.Printf("2**%d\n", i1 )
//	}
//
//	fmt.Printf("__________________\n" )
//
//	for _,v1 := range pow {
//	    fmt.Printf("2**%d\n", v1 )
//	}
//
//	fmt.Printf("__________________\n" )
//
//}
// index 41  switch
//
//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func main() {
//	fmt.Print("Go runs on ")
//	switch os := runtime.GOOS; os {
//	case "darwin" :
//		fmt.Println("OS X.")
//
//	case "linux":
//		fmt.Println("Linux.")
//	default:
//		// freebsd, openbsd,
//		// plan9, windows...
//		fmt.Printf("%s.", os)
//	}
//}

//
//
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	fmt.Println("When's Saturday?")
//	today := time.Now().Weekday()
//	switch time.Saturday {
//	case today+0:
//		fmt.Println("Today.")
//
//	case today+1:
//		fmt.Println("Tomorrow.")
//	case today+2:
//		fmt.Println("In two days.")
//	default:
//		fmt.Println("Too far away.", " time.Saturday=", time.Saturday)
//	}
//}
//
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	t := time.Now()
//
//	switch {
//	case t.Hour() < 12:
//		fmt.Println("< 12")
//		fallthrough
//	case t.Hour() < 17:
//		fmt.Println(" < 17")
//		fallthrough
//	case t.Hour() < 19:
//		fmt.Println(" < 19")
//	default:
//		fmt.Println("Good evening.")
//	}
//}

//package main
//
//import "fmt"
//
//// fibonacci 函数会返回一个返回 int 的函数。
//func fibonacci() func() int {
//	x := 0
//	y := 1
//	z := 0
//	return func() int {
//		z = x+y
//		x = y
//		y = z
//		return z
//	}
//}
//
//func main() {
//	f := fibonacci()
//	for i := 0; i < 10; i++ {
//		fmt.Println(f())
//	}
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	v := Vertex{3, 4}
//	fmt.Println(v.Abs())
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type MyFloat float64
//
//func (f * MyFloat) Abs() float64 {
//	if *f < 0 {
//		return float64(-*f)
//	}
//	return float64(*f)
//}
//
//func main() {
//	f := MyFloat(-math.Sqrt2)
//	fmt.Println(f.Abs())
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)

//type Vertex struct {
//	X, Y float64
//}
//
//func (v Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	v := &Vertex{3, 4}
//	v.Scale(5)
//	fmt.Println(v, v.Abs())
//}

//
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//type Abser interface {
//	Abs() float64
//}
//
//func main() {
//	var a Abser
//	f := MyFloat(-math.Sqrt2)
//	v := Vertex{3, 4}
//
//	a = f  // a MyFloat implements Abser
//
//	fmt.Println(a.Abs())
//
//	a = &v // a *Vertex implements Abser
//
//	fmt.Println(a.Abs())
//
////	a = v  // a Vertex, does NOT
//	       // implement Abser
//
//}
//
//type MyFloat float64
//
//func (f MyFloat) Abs() float64 {
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

//
//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//type Reader interface {
//	Read(b []byte) (n int, err error)
//}
//
//type Writer interface {
//	Write(b []byte) (n int, err error)
//}
//
//type ReadWriter interface {
//	Reader
//	Writer
//}
//
//func main() {
//	var w Writer
//
//	// os.Stdout implements Writer
//	w = os.Stdout
//
//	fmt.Fprintf(w, "hello, writer\n")
//}

//index 55  Error

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//type MyError struct {
//	When time.Time
//	What string
//}
//
//func (e *MyError) Error() string {
//	return fmt.Sprintf("at %v, %s",
//		e.When, e.What)
//}
//
//func run() error {
//	return &MyError{
//		time.Now(),
//		"it didn't work",
//	}
//}
//
//func main() {
//	if err := run(); err != nil {
//		fmt.Println(err)
//	}
//}

//
//package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//type Hello struct{}
//
//func (h Hello) ServeHTTP(
//		w http.ResponseWriter,
//		r *http.Request) {
//
//	fmt.Println( "req=", r  )
//
//	fmt.Println( "end of req_________________"  )
//	fmt.Fprint(w, "Hello!")
//}
//
//func main() {
//	var h Hello
////	http.ListenAndServe("localhost:4000",h)
//	http.ListenAndServe("192.168.9.26:4000",h)
//}
//
//
//package main
//
//import (
//	"fmt"
//	"runtime"
//	"time"
//)
//
//func say(s string) {
//	for i := 0; i < 5; i++ {
//		runtime.Gosched()		//让出处理器
//		fmt.Println(s)
//	}
//}
//
//func main() {
//	//	go say("world")
//	say("world")
//	say("hello")
//	time.Sleep(100 * 1e9)
//}

//package main
//
//import "fmt"
//
//func sum(a []int, c chan int) {
//	sum := 0
//	for _, v := range a {
//		sum += v
//	}
//	c <- sum // send sum to c
//}
//
//func main() {
//	a := []int{7, 2, 8, -9, 4, 0}
//
//	c := make(chan int)
//	go sum(a[:len(a)/2], c)
//	go sum(a[len(a)/2:], c)
//	x, y := <-c, <-c // receive from c
//
//	fmt.Println(x, y, x+y)
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	c := make(chan int, 2)
//	c <- 1
//	c <- 2
//	c <- 3
//	c <- 4
//	fmt.Println(<-c)
//	fmt.Println(<-c)
//}

//index 65  range &  close
//package main
//
//import (
//	"fmt"
//)
//
//func fibonacci(n int, c chan int) {
//        x, y := 1, 1
//        for i := 0; i < n; i++ {
//                c <- x
//                x, y = y, x + y
//        }
//        close(c)
//}
//
//func main() {
//        c := make(chan int, 10)
//	go fibonacci(cap(c), c)
//        for i := range c {
//                fmt.Println(i)
//        }
//}
//
//package main
//
//import (
//	"fmt"
//)
//
//func fibonacci(c chan int) {
//	x, y := 1, 1
//	for i := 0; i < 10; i++ {
//		c <- x
//		x, y = y, x+y
//	}
//}
//
//func main() {
//	c := make(chan int)
//	go fibonacci(c)
//	for i := range c {
//		fmt.Println(i)
//	}
//}

//package main
//
//import "fmt"
//
//func fibonacci(c, quit chan int) {
//	x, y := 1, 1
//	for {
//		select {
//		case c <- x:
//			x, y = y, x+y
//		case <-quit:
//			fmt.Println("quit")
//			return
//		}
//	}
//}
//
//func main() {
//	c := make(chan int)
//	quit := make(chan int)
//	go func() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(<-c)
//		}
//		quit <- 0
//	}()
//	fibonacci(c, quit)
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	tick := time.Tick(1e8)
//	boom := time.After(5e8)
//	for {
//		select {
//		case i := <-tick:
//			fmt.Println("tick.", "i=", i)
//		case i := <-boom:
//			fmt.Println("BOOM.", "i=", i)
//			return
//		default:
//			fmt.Println("    .")
//			time.Sleep(5e7)
//		}
//	}
//}
//
//package main
//
//import (
//	"fmt"
//)
//
//type Fetcher interface {
//        // Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
//	Fetch(url string) (body string, urls []string, err error)
//}
//
//// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
//func Crawl(url string, depth int, fetcher Fetcher) {
//        // TODO: 并行的抓取 URL。
//        // TODO: 不重复抓取页面。
//        // 下面并没有实现上面两种情况：
//	if depth <= 0 {
//		return
//	}
//	body, urls, err := fetcher.Fetch(url)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Printf("found: %s %q\n", url, body)
//	for _, u := range urls {
//		Crawl(u, depth-1, fetcher)
//	}
//	return
//}
//
//func main() {
//	Crawl("http://www.sina.com.cn/", 4, fetcher)
//}
//
//
//// fakeFetcher 是返回若干结果的 Fetcher。
//type fakeFetcher map[string]*fakeResult
//
//type fakeResult struct {
//	body string
//	urls     []string
//}
//
//func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
//	if res, ok := (*f)[url]; ok {
//		return res.body, res.urls, nil
//	}
//	return "", nil, fmt.Errorf("not found: %s", url)
//}
//
//// fetcher 是填充后的 fakeFetcher。
//var fetcher = &fakeFetcher{
//	"http://golang.org/": &fakeResult{
//		"The Go Programming Language",
//		[]string{
//			"http://golang.org/pkg/",
//			"http://golang.org/cmd/",
//		},
//	},
//	"http://golang.org/pkg/": &fakeResult{
//		"Packages",
//		[]string{
//			"http://golang.org/",
//			"http://golang.org/cmd/",
//			"http://golang.org/pkg/fmt/",
//			"http://golang.org/pkg/os/",
//		},
//	},
//	"http://golang.org/pkg/fmt/": &fakeResult{
//		"Package fmt",
//		[]string{
//			"http://golang.org/",
//			"http://golang.org/pkg/",
//		},
//	},
//	"http://golang.org/pkg/os/": &fakeResult{
//		"Package os",
//		[]string{
//			"http://golang.org/",
//			"http://golang.org/pkg/",
//		},
//	},
//}

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package main

// import (
// 	"log"
// 	"net/http"
// 	"time"
// )

// const (
// 	numPollers     = 2                // number of Poller goroutines to launch
// 	pollInterval   = 60 * time.Second // how often to poll each URL
// 	statusInterval = 10 * time.Second // how often to log status to stdout
// 	errTimeout     = 10 * time.Second // back-off timeout on error
// )

// var urls = []string{
// 	//	"http://www.google.com/",
// 	//	"http://golang.org/",
// 	//	"http://blog.golang.org/",
// 	//	"http://www.sina.com.cn",
// 	//	"http://www.10086.com",
// 	"http://123",
// }

// // State represents the last-known state of a URL.
// // url 的状态
// type State struct {
// 	url    string
// 	status string
// }

// // StateMonitor maintains a map that stores the state of the URLs being
// // polled, and prints the current state every updateInterval nanoseconds.
// // It returns a chan State to which resource state should be sent.
// func StateMonitor(updateInterval time.Duration) chan<- State {
// 	updates := make(chan State)
// 	urlStatus := make(map[string]string)
// 	ticker := time.NewTicker(updateInterval)
// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				logState(urlStatus)
// 			case s := <-updates:
// 				urlStatus[s.url] = s.status
// 			}
// 		}
// 	}()
// 	return updates
// }

// //	@brief 打印url的连接状态
// // 		logState prints a state map.
// func logState(s map[string]string) {
// 	log.Println("Current state:")
// 	for k, v := range s {
// 		log.Printf(" %s %s", k, v)
// 	}
// }

// // @brief 资源   一个资源表示一个 url
// // 		Resource represents an HTTP URL to be polled by this program.
// type Resource struct {
// 	url      string
// 	errCount int
// }

// // @brief 执行http HEAD 请求
// // 		Poll executes an HTTP HEAD request for url
// //		and returns the HTTP status string or an error string.
// func (r *Resource) Poll() string {
// 	resp, err := http.Head(r.url)

// 	log.Println("111111111")

// 	if err != nil {
// 		log.Println("Error", r.url, err)
// 		r.errCount++
// 		log.Println("  r.errCount=", r.errCount)
// 		return err.Error()
// 	}
// 	r.errCount = 0
// 	return resp.Status
// }

// // @brief 将url放回请求队列
// //		Sleep sleeps for an appropriate interval (dependant on error state)
// //		before sending the Resource to done.
// func (r *Resource) Sleep(done chan<- *Resource) {
// 	time.Sleep(pollInterval + errTimeout*time.Duration(r.errCount))
// 	done <- r
// }

// // @brief 业务处理函数
// //		从url请求队列获取url， 执行url访问操作,再将url 放回 url完成队列
// func Poller(in <-chan *Resource, out chan<- *Resource, status chan<- State) {
// 	for r := range in {
// 		s := r.Poll()
// 		status <- State{r.url, s}
// 		out <- r
// 	}
// }

// func main() {
// 	//初始化url请求队列, url完成队列
// 	//	Create our input and output channels.
// 	pending, complete := make(chan *Resource), make(chan *Resource)

// 	//初始化监控\打印 go程
// 	//	Launch the StateMonitor.
// 	status := StateMonitor(statusInterval)

// 	// Launch some Poller goroutines.
// 	for i := 0; i < numPollers; i++ {
// 		go Poller(pending, complete, status)
// 	}

// 	// Send some Resources to the pending queue.
// 	go func() {
// 		for _, url := range urls {
// 			pending <- &Resource{url: url}
// 		}
// 	}()

// 	for r := range complete {
// 		go r.Sleep(pending)
// 	}
// }
//
