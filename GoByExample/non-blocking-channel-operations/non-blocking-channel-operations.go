// Basic sends and receives on channels are blocking.
// However, we can use `select` with a `default` clause to
// implement _non-blocking_ sends, receives, and even
// non-blocking multi-way `select`s.

package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here's a non-blocking receive. If a value is
	// available on `messages` then `select` will take
	// the `<-messages` `case` with that value. If not
	// it will immediately take the `default` case.

	// 非阻塞的 channel receive
	// "messages" 有值则直接读出， 否则将不阻塞在该 channel中， 直接执行 default
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send works similarly.
	// 非阻塞的 channel send
	// 不会往 "messages" 写入数据，而是直接执行 default
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 另起 go routine， 10秒后  异步发送 “world”
	go func() {
		time.Sleep(10 * time.Second)
		msg = "world"
		messages <- msg
	}()

	// We can use multiple `case`s above the `default`
	// clause to implement a multi-way non-blocking
	// select. Here we attempt non-blocking receives
	// on both `messages` and `signals`.
	i := 0

	// 非阻塞的 channel receive
	// "messages" 有值则直接读出， 否则将不阻塞在该 channel中， 直接执行 default
	// 10秒后 "messages" 有值读出，此时 default 已经执行了 18M 次以上（视应用所在CPU的执行能力）
	for {
		i = i + 1
		select {
		case msg := <-messages:
			fmt.Println("received message|msg=", msg, "|i=", i)
			break
		case sig := <-signals:
			fmt.Println("received signal", sig)
		default:
			//fmt.Println("no activity,i=%v", i)
		}
	}

}
