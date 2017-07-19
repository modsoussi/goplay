package main

import "fmt"

func writeTo(c chan int, n int) {
	c <- n * n
}

func main() {
	c := make(chan int)
	go writeTo(c, 2)
	fmt.Println("Waiting ...")
	go writeTo(c, 3)
	i := <-c
	fmt.Println(i)
	j := <-c
	fmt.Println(j)
	go writeTo(c, 8)
	k := <-c
	fmt.Println(k)
}
