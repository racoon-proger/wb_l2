package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	// Парсинг аргументов командной строки
	host := flag.String("host", "", "IP адрес или доменное имя сервера")
	port := flag.String("port", "", "Порт сервера")
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут подключения")
	flag.Parse()

	if *host == "" || *port == "" {
		flag.Usage()
		return
	}
	address := fmt.Sprintf("%s:%s", *host, *port)

	// Подключение к серверу
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Println("Ошибка подключения к серверу:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Подключено к", conn.RemoteAddr())

	// Перехват сигнала выхода
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go readData(conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		// Чтение данных из STDIN
		line, err := reader.ReadString('\n')
		if err != nil {
			// Программа завершается при ошибке чтения данных из STDIN
			fmt.Println("Ошибка чтения данных из STDIN:", err)
			break
		}
		// Удаление символа новой строки из введенной строки
		text := strings.TrimSuffix(line, "\n")
		// Отправка данных в сокет
		_, err = conn.Write([]byte(text))
		if err != nil {
			// Программа завершается при ошибке записи в сокет
			fmt.Println("Ошибка записи данных в сокет:", err)
			break
		}
		// Ожидание сигнала о выходе
		select {
		case <-signalChan:
			fmt.Println("Выход...")
			return
		default:
		}
	}
}

// readData - функция для чтения данных из сокета и вывода их в STDOUT
func readData(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		// Чтение данных из сокета
		line, err := reader.ReadString('\n')
		if err != nil {
			// Программа завершается при ошибке чтения данных из сокета
			fmt.Println("Ошибка чтения данных из сокета:", err)
			os.Exit(1)
		}
		// Вывод данных в STDOUT
		fmt.Print(line)
	}
}
