package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		// Чтение команды пользователя
		scanner.Scan()
		command := scanner.Text()
		// Разделение команд на отдельные части
		parts := strings.Split(command, " ")
		cmd := parts[0]
		args := parts[1:]
		// Обработка команды в зависимости от ее типа
		switch cmd {
		case "cd":
			// Смена директории
			err := os.Chdir(args[0])
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "pwd":
			// Вывод текущего пути
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(cwd)
			}
		case "echo":
			// Вывод аргумента в STDOUT
			fmt.Println(strings.Join(args, " "))
		case "kill":
			// "Убийство" процесса
			err := exec.Command("kill", args...).Run()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "ps":
			// Вывод информации о запущенных процессах
			cmd := exec.Command("ps")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
			}
		default:
			// Обработка остальных команд
			cmd := exec.Command(parts[0], parts[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}
