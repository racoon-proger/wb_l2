package main

func main() {
	// создаем канал
	ch := make(chan int)
	// запускаем горутину с циклом
	go func() {
		// в каждой итерации цикла данные записываются в канал
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	// с помощью range читаем данные из канала
	for n := range ch {
		println(n)
	}
}

/* вывод:
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!
main зависнет т.к как канал не закрыт
*/
