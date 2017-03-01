// In a [previous](range) example we saw how `for` and
// `range` provide iteration over basic data structures.
// We can also use this syntax to iterate over
// values received from a channel.

package main

import (
	"fmt"
	"time"
)

func main() {

	// We'll iterate over 2 values in the `queue` channel.
	queue := make(chan string, 2)

	// go routine， channel push 2个数据;
	// close(channe1) 保证后续的 range 在遍历所有元素后 不死锁
	go func() {
		//push
		queue <- "one"

		//睡眠1秒
		time.Sleep(time.Second * 1)

		//push
		queue <- "two"

		//睡眠1秒
		time.Sleep(time.Second * 1)

		//关闭
		close(queue)
	}()

	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.

	// for..range channel， 持续从 channel 读取数据
	//	当 channel 关闭后，该 for..range 语句会 break
	for elem := range queue {
		fmt.Println(elem)
	}

	//睡眠N秒
	// 使得以上异步操作得以全部实行
	time.Sleep(time.Second * 5)
	fmt.Println("app end")
}
