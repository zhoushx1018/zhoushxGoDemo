// A _goroutine_ is a lightweight thread of execution.

package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.
	f("direct")

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	go f("goroutine")

	// You can also start a goroutine for an anonymous
	// function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now, so execution falls through
	// to here. This `Scanln` code requires we press a key
	// before the program exits.

	// 以上 go routine 是异步执行的，因此在这里做个  scanln ，以等待以上 go routine 执行完成
	// 	在命令行输入  "done" 后， 进程才会退出
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}