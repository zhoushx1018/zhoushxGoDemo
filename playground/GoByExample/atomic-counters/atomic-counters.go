//使用 sync/atomic 做原子操作

package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			//for {

				atomic.AddUint64(&ops, 1)

				time.Sleep(time.Millisecond)
			//}
		}()
	}

	// Wait a second to allow some ops to accumulate.

	//睡眠
	time.Sleep(time.Second)

	//获取当前 累加结果
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}