package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// wg := &sync.WaitGroup{}

	// // สร้างเท่าจำนวนที่เราต้องการ wait
	// wg.Add(3)

	// start := time.Now()
	// go doSomething(wg)
	// go doSomething(wg)
	// go doSomething(wg)

	// // รอจนกว่า goroutine ทุกตัวจะทำงานเสร็จ
	// wg.Wait()
	// fmt.Println(time.Since(start))
	// example()

	// race condition example
	// loop()

	ch := make(chan int)

	// ต้องการส่ง signal เฉย ๆ จึงเลือกใช้ type ที่ไม่กิน memory มาก
	qCh := make(chan struct{})

	go fibonacci(ch, qCh)

	for i := 0; i < 10; i++ {
		// ch จะรอจนกว่าจะมีของส่งเข้ามาแล้วจึง print
		fmt.Println(<-ch)
	}

	qCh <- struct{}{}

	<-qCh
	fmt.Println("end")
}

func doSomething(wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("something")
	wg.Done()
}

func example() {
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("-")
		}
	}()
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("+")
		}
	}()
	time.Sleep(time.Second)
}

var i int
var mux sync.Mutex

type thing struct {
	i   int
	mux sync.Mutex
}

func loop() {
	go func() {
		for {
			fmt.Println(read())
		}
	}()

	for i := 0; i < 10; i++ {
		write(i)
	}
}

func write(n int) {
	mux.Lock()
	i = n
	mux.Unlock()
}

func read() int {
	mux.Lock()
	defer mux.Unlock()
	return i
}

func fibonacci(ch chan int, qCh chan struct{}) {
	a, b := 0, 1

	for {
		select {
		case ch <- a:
			a, b = b, a+b
		case <-qCh:
			qCh <- struct{}{}
			return
		}
	}
}
