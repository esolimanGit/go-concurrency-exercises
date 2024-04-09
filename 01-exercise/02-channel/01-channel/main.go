package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)
	c := <-ch
	fmt.Printf("computed value %v\n", c)
}
