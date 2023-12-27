package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	uniteChannels()
}

func uniteChannels() {
	or := func(channels ...<-chan interface{}) <-chan interface{} {
		out := make(chan interface{})
		wg := sync.WaitGroup{}
		wg.Add(len(channels))

		for _, channel := range channels {
			// запускаем горутину которая принимает канал
			// читает из него и пишет в другой канал
			go func(ch <-chan interface{}) {
				for value := range ch {
					out <- value
				}
				fmt.Printf("Channel closed right now\n")
				wg.Done()
			}(channel)
		}
		go func() {
			wg.Wait()
			close(out)
		}()
		return out
	}
	sig := func(after time.Duration) <-chan interface{} {
		// создаем канал, в который можно писать и читать
		c := make(chan interface{})
		go func() {
			defer close(c)
			c <- "some value"
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	output := or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
	)
	for v := range output {
		fmt.Println(v)
	}
	fmt.Println(time.Since(start))
}
