package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

/* вывод:
цикл будет бесконечным, мы постоянно будем читать данные из канала
потому что в конструкции select не проверяем закрытие канала с которого читаем.
лучше воспользоваться циклом range
go func() {
		for v := range a {
			c <- v
		}
	}()
	go func() {
		for v := range b {
			c <- v
		}
	}()
	либо добовлять вторую булевую переменную в case что бы проверять
	закрытие канала
*/
