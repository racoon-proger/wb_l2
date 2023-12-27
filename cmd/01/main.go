package main

import (
	"log"
	"os"

	"github.com/beevik/ntp"
)

func timeOutput() {
	// вывод в stderr
	logger := log.New(os.Stderr, "", 0)
	// запрос текущего времени
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		logger.Fatal(err.Error())
	}
	log.Printf("The current time is %s", time)

}

func main() {
	timeOutput()
}
