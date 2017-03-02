package main

//本例子是自己写的，不是 GoByExample 中的例子
//  GoByExample 的channel 例子，很容易死锁 (fatal error: all goroutines are asleep - deadlock!)
//  本例子演示，如何写出不死锁（fatal error: all goroutines are asleep - deadlock!) 的程序

import (
	"fmt"
	"time"
)

func pushChan(cCurrChan chan<- int) {

	iTmp := 0
	for {
		time.Sleep(time.Second * 1)
		cCurrChan <- iTmp
		iTmp = iTmp + 1
	}
}

func main() {
	cTmp := make(chan int)
	go pushChan(cTmp)

	for {
		iTmp := <-cTmp
		fmt.Println("iTmp=", iTmp)
	}

}
