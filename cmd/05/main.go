package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	afterFlag := flag.Int("A", 0, "печатать +N строк после совпадения")
	beforeFlag := flag.Int("B", 0, "печатать +N строк до совпадения")
	contextFlag := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	countFlag := flag.Bool("c", false, "количество строк")
	ignoreCaseFlag := flag.Bool("i", false, "игнорировать регистр")
	invertFlag := flag.Bool("v", false, "вместо совпадения, исключать")
	fixedFlag := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNumFlag := flag.Bool("n", false, "напечатать номер строкиr")
	flag.Parse()

	// получаем имя файла
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("пустое имя файла")
		return
	}
	filename := args[0]

	// читаем содержимое файла
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var lines []string
	for _, line := range strings.Split(string(content), "\n") {
		if line != "" {
			lines = append(lines, line)
		}
	}
	// номер строки
	lineNum := 0
	// кол-во совпадений
	matchCount := 0
	// строка до совпадения
	beforeLines := []string{}
	// строка после совпадения
	afterLines := []string{}

	// перебираем строки
	for _, line := range lines {
		lineNum++
		match := false
		if *fixedFlag {
			match = line == filename
		} else {
			if *ignoreCaseFlag {
				match = strings.Contains(strings.ToLower(line), strings.ToLower(filename))
			} else {
				match = strings.Contains(line, filename)
			}
		}
		if (*invertFlag && !match) || (!*invertFlag && match) {
			matchCount++
			// вывести номер строки: flag -n
			if *lineNumFlag {
				fmt.Printf("%d:", lineNum)
			}
			fmt.Println(line)
			// вывести строки до совпадения: flag -b
			for _, beforeLine := range beforeLines {
				fmt.Println(beforeLine)
			}
			// вывести строки после совпадения: flag -a
			for _, afterLine := range afterLines {
				fmt.Println(afterLine)
			}
		} else if *contextFlag > 0 && matchCount > 0 && matchCount <= *contextFlag {
			// строки вокруг совпадения: flag -C
			fmt.Println(line)
		}
		beforeLines = append(beforeLines, line)
		if len(beforeLines) > *beforeFlag {
			beforeLines = beforeLines[1:]
		}
		afterLines = append(afterLines, line)
		if len(afterLines) > *afterFlag {
			afterLines = afterLines[1:]
		}
	}
	// вывести кол-во строк: flag -c
	if *countFlag {
		fmt.Println("количество строк", matchCount)
	}
}
