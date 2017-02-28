// _Closing_ a channel indicates that no more values
// will be sent on it. This can be useful to communicate
// completion to the channel's receivers.

package main

import "fmt"

// In this example we'll use a `jobs` channel to
// communicate work to be done from the `main()` goroutine
// to a worker goroutine. When we have no more jobs for
// the worker we'll `close` the `jobs` channel.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Here's the worker goroutine. It repeatedly receives
	// from `jobs` with `j, more := <-jobs`. In this
	// special 2-value form of receive, the `more` value
	// will be `false` if `jobs` has been `close`d and all
	// values in the channel have already been received.
	// We use this to notify on `done` when we've worked
	// all our jobs.

	// 异步 go routine
	// 从阻塞 channel  获取数据，获取到数据时 , "more" 为 true
	// 当阻塞 channel 已经被关闭时， "more" 为false
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				// more 为 false， 说明 channel 已经被关闭, 其中的数据已经全部被读出
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

	// We await the worker using the
	// [synchronization](channel-synchronization) approach
	// we saw earlier.

	//等待异步 go routine 的退出;
	//	否则main  routine 先退出，无法完整的观察整个 go routine 的执行过程
	<-done
}