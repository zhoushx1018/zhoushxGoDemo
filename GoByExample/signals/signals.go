// Sometimes we'd like our Go programs to intelligently
// handle [Unix signals](http://en.wikipedia.org/wiki/Unix_signal).
// For example, we might want a server to gracefully
// shutdown when it receives a `SIGTERM`, or a command-line
// tool to stop processing input if it receives a `SIGINT`.
// Here's how to handle signals in Go with channels.

//本例子最好在 linux 下执行
// 在 windows 下执行不太正常（估计是Windows 对 posix 信号 支持的不是很好）

//	[zhoushx@dev926 signals]$ ./signals
//	[signals.go 48]  awaiting signal
//	[signals.go 41]  sig= interrupt
//	[signals.go 51]  sleep 5 sec
//	[signals.go 53]  exiting

package main

import "fmt"
import "os"
import "os/signal"
import (
	"syscall"
	"time"
)

import . "zhoushxGoDemo/GoByExample/fileline"

func main() {

	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications (we'll also make one to
	// notify us when the program can exit).
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.

	//在 channel 中登记 相关 信号
	//	一旦进程收到相关信号， 内核将会将相关信号 push 到 channel中
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.

	//goroutine 阻塞等待 channel
	go func() {
		sig := <-sigs
		fmt.Println(FileLine(), "sig=", sig)
		done <- true
	}()

	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.
	fmt.Println(FileLine(), "awaiting signal")
	<-done

	fmt.Println(FileLine(), "sleep 5 sec")
	time.Sleep(time.Second * 5)
	fmt.Println(FileLine(), "exiting")
}
