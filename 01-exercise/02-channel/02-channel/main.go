package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	// TODO: range over channel to recv values
	for v := range ch {
		fmt.Println(v)
	}
}
