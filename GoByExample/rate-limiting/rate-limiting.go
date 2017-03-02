// _[Rate limiting](http://en.wikipedia.org/wiki/Rate_limiting)_
// is an important mechanism for controlling resource
// utilization and maintaining quality of service. Go
// elegantly supports rate limiting with goroutines,
// channels, and [tickers](tickers).

package main

import "time"
import "fmt"

func main() {

	// First we'll look at basic rate limiting. Suppose
	// we want to limit our handling of incoming requests.
	// We'll serve these requests off a channel of the
	// same name.

	// 向chan 连续push 5个元素
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This `limiter` channel will receive a value
	// every 200 milliseconds. This is the regulator in
	// our rate limiting scheme.

	//新建 ticker， 每 200毫秒 tick 一次
	limiter := time.Tick(time.Millisecond * 200)

	// By blocking on a receive from the `limiter` channel
	// before serving each request, we limit ourselves to
	// 1 request every 200 milliseconds.

	//通过使用  ticker，实现 每 200毫秒 pop一次 chan
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//-------------------------------------------------------------------------------
	// We may want to allow short bursts of requests in
	// our rate limiting scheme while preserving the
	// overall rate limit. We can accomplish this by
	// buffering our limiter channel. This `burstyLimiter`
	// channel will allow bursts of up to 3 events.

	//新建 cap 为3 的chan
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 milliseconds we'll try to add a new
	// value to `burstyLimiter`, up to its limit of 3.

	//每隔 200毫秒 push 一次 chan
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incoming requests. The first
	// 3 of these will benefit from the burst capability
	// of `burstyLimiter`.

	//新建 chan， push 5个元素
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// chan pop5次， 前三次pop是立即执行的， 后两次pop 有200毫秒的间隔
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}