package main

import (
	"errors"
	"fmt"
	"time"
)

func waiter(d time.Duration) {
	time.Sleep(d)
	fmt.Println("Waiter: -Hello, World!")
}

func waitergives(s string, c chan string, d time.Duration) {
	time.Sleep(d)
	c <- s
}

func tupleReturn() (string, int, error) {
	return "Hi", 1, errors.New("error")
}

func immediate() {
	fmt.Println("Immediate: -Hello, World!")
}

func main() {
	duration, _ := time.ParseDuration("2s")
	go waiter(duration)
	go immediate()

	c := make(chan string)
	duration, _ = time.ParseDuration("4s")
	go waitergives("Well fuck!", c, duration)
	s := <-c
	fmt.Println(s)

	a, b, err := tupleReturn()
	fmt.Println(a, b, err)
}
