package main

import "time"

func main() {

	channel := make(chan string)

	go func() {
		channel <- "Hello World"
	}()

	//Main Thread
	result := <-channel
	println(result)

	//Threads in Go are called goroutines
	//go counter("a")
	//go counter("b")

	//time.Sleep(11 * time.Second)
}

func counter(label string) {
	for i := 0; i < 10; i++ {
		println(label, i)
		time.Sleep(1 * time.Second)
	}
}
