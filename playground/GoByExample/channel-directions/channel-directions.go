//当使用channel 作为功能参数时，可以指定是仅发送还是接收
// 这种特性增加了类型安全性

package main

import "fmt"

// ping 函数仅允许  channel for sending 的形参
//   如果使用 channel for receive 的实参，则会产生  编译器的错误
//   从而规避实参传错的风险
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong 函数指明 第一、二个形参，分别为  channel for sending 和 channel for receive
//    规避实参传错的风险
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Printf("pongs=%s\n",<-pongs)
}