// channel 默认是非缓冲的

package main

import "fmt"

func main() {

	// 缓冲区大小为2的channel
	messages := make(chan string, 2)

	//注意，此处没有用  go routine
	//  由于 channel 是有2个缓存单元，因此以下两个 push， 直接把 channel 的两个缓存单元 push满
	//  此时不会阻塞（死锁）
	messages <- "buffered"
	messages <- "channel"

	//缓冲区满了，如果继续push ，应用会阻塞(死锁)
	//messages <- "go"

	// Later we can receive these two values as usual.
	//从 channel 取出 2个元素
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}