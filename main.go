package main

import (
	"fmt"
	"time"
)

const reset = "\033[0m"
const green = "\033[32m"
const yellow = "\033[33m"

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- green + "ping" + reset
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- yellow + "pong" + reset
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go ping(c)
	go pong(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
