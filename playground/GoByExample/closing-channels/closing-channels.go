package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 异步 go routine
	// 从阻塞 channel  获取数据，获取到数据时 , "ok" 为 true
	// 当阻塞 channel 已经被关闭时， "ok" 为false
	go func() {
		for {
			j, ok := <-jobs
			if ok {
				fmt.Println("received job", j)
			} else {
				// ok 为 false， 说明 channel 已经被关闭, 其中的数据已经全部被读出
				//	此时可以结束读取数据
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// This sends 3 jobs to the worker over the `jobs`
	// channel, then closes it.

	//发送 3个数据 给 channel
	//  发送后，将 channel 关闭
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")


	//等待异步 go routine 的退出;
	//	否则main  routine 先退出，无法完整的观察整个 go routine 的执行过程
	<-done
}