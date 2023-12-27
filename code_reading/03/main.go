package main

import (
	"fmt"
	"os"
)

// Foo возвращает интерфейс error которому присваевается
// переменная err типа os.PathError
func Foo() error {
	// создаем переменную типа структуры os.PathError которая равна nil
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}

/*
когда интерфейсу который равен nil присваеваем любой тип даже если присваеваемый
тип nil интерфейс уже не равен nil
*/
