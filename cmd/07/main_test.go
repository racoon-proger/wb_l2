package main

import (
	"testing"
	"time"
)

func TestTaskSeven(t *testing.T) {
	// Создаю каналы
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	// Создаю горутины, которые будут писать значения в каналы
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "value 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "value 2"
	}()

	// Вызываю функцию TaskSeven()
	TaskSeven(ch1, ch2)

	// Ожидаемые значения
	expected := []string{"value 2", "value 1"}

	// Проверяю полученные значения
	for i, v := range expected {
		if v != output[i] {
			t.Errorf("Ожидается %s, но получено %s", v, output[i])
		}
	}
}
