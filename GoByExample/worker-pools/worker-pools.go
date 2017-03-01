// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import "fmt"
import "time"

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.

// worker 的业务逻辑
// 从 jobs 持续（阻塞）读取， 处理数据（sleep，把相关整数翻倍）后 ， 将结果push 给 results
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.

	//生成3个worker的 go routine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.

	//生产 5个int， push 给 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)  //关闭 jobs，否则会报死锁 错误

	 //Finally we collect all the results of the work.
	 //等待 results 结果的输出
	//for a := 1; a <= 5; a++ {
	for a := 1; ; a++ {
		<-results
	}



	//i := 0
	//for a := range results {
	//	i = i + 1
	//	fmt.Println("i:", i, "result:", a)
	//
	//}


}
