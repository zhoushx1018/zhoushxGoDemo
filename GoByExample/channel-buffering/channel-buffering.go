// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.

package main

import "fmt"

func main() {

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.

	//注意，此处没有用  go routine
	//  由于 channel 是有2个缓存单元，因此以下两个 push， 直接把 channel 的两个缓存单元 push满
	//  此时不会阻塞（死锁）
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	//从 channel 取出 2个元素
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}