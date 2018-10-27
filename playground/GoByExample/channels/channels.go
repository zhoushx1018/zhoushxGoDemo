package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建 channel  和 匿名携程
	//  主携程（主进程） 和 匿名携程之间使用 channel 通信
	messages := make(chan string)

	go func() {
		time.Sleep(5*time.Second)		//睡眠5秒
		messages <- "ping"
		}()

	//等待 channel 有数据可读
	msg := <-messages
	fmt.Println(msg)
}