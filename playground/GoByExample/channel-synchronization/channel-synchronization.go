package main

import "fmt"
import "time"

// 睡眠5秒 再 push  channel
func worker(done chan bool) {
	fmt.Print("working...\n")
	time.Sleep(5*time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	// 创建 channel  和 worker携程
	//  主携程（主进程） 和 worker携程之间使用 channel 通信
	done := make(chan bool, 1)
	go worker(done)

	//等待 channel 有数据可读
	<-done
}