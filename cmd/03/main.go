package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortFile(
	filename string,
	columnNumber int,
	numericFlag, reverseFlag, uniqueFlag bool,
) (
	[]string,
	error,
) {
	// читаем содержимое файла
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	// собираем все строки и откидываем пустые
	var lines []string
	for _, line := range strings.Split(string(content), "\n") {
		if line != "" {
			lines = append(lines, line)
		}
	}
	sort.Slice(lines, func(i, j int) bool {
		fieldsI := strings.Fields(lines[i])
		fieldsJ := strings.Fields(lines[j])

		if columnNumber > 0 && columnNumber <= len(fieldsI) && columnNumber <= len(fieldsJ) {
			fieldsI = fieldsI[columnNumber-1:]
			fieldsJ = fieldsJ[columnNumber-1:]
		}
		// сортировка по числовому значению: flag -n
		if numericFlag {
			numI, errI := strconv.Atoi(fieldsI[0])
			numJ, errJ := strconv.Atoi(fieldsJ[0])
			if errI == nil && errJ == nil {
				fieldsI[0] = strconv.Itoa(numI)
				fieldsJ[0] = strconv.Itoa(numJ)
			}
		}
		less := strings.Join(fieldsI, " ") < strings.Join(fieldsJ, " ")
		// сортировка в обратном порядке: flag -r
		if reverseFlag {
			return !less
		}
		return less
	})

	// удаляем повторяющиеся строки: flag -u
	if uniqueFlag {
		uniqueLines := make(map[string]struct{})
		var result []string
		for _, line := range lines {
			if _, ok := uniqueLines[line]; !ok {
				uniqueLines[line] = struct{}{}
				result = append(result, line)
			}
		}
		lines = result
	}
	return lines, nil
}

func main() {
	columnNumber := flag.Int("k", 0, "сортировать по колонке")
	numericFlag := flag.Bool("n", false, "строгая сортировка по числам")
	reverseFlag := flag.Bool("r", false, "обратная сортировка")
	uniqueFlag := flag.Bool("u", false, "отбросить дубликаты")
	flag.Parse()

	// получаем имя файла
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("не указано имя файла")
		return
	}
	filename := args[0]

	output, err := sortFile(filename, *columnNumber, *numericFlag, *reverseFlag, *uniqueFlag)
	if err != nil {
		fmt.Println(err)
		return
	}
	// выводим результат
	for _, line := range output {
		fmt.Println(line)
	}
}
