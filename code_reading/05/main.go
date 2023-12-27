package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// do something
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

/* вывод: error
   потому что мы присвоили интерфейсу error структуру customError после чего
   интерфейс error уже не был nil, хотя если сравнить интерфейс error
   путем err == nil будет true
*/
