package main

import (
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	anagrams := searchAnagrams([]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"})
	log.Println(anagrams)
}

func anagrams(word string, dict map[int]string) []string {
	validWord := strings.ToLower(strings.TrimSpace(word))
	validWordLen := utf8.RuneCountInString(validWord)
	// Формируем регулярное выражение вида []{}
	var regExpr strings.Builder
	regExpr.WriteString("[")
	for _, symbol := range validWord {
		regExpr.WriteString(string(symbol) + ",")
	}
	regExpr.WriteString("]{" + strconv.Itoa(validWordLen) + "}")
	// поиск анаграм с помощью регулярных выражений
	re := regexp.MustCompile(regExpr.String())
	return re.FindAllString(dict[validWordLen], -1)
}

func searchAnagrams(words []string) map[string][]string {
	dict := make(map[int]string)
	for _, word := range words {
		// в ключ кладем число символов в слове
		// в значение кладем слово в нижнем регистре
		dict[utf8.RuneCountInString(word)] += strings.ToLower(strings.TrimSpace(word)) + " "
	}
	// wordList - список слов для поиска их анаграм
	wordList := []string{"столик"}
	result := make(map[string][]string)
	for _, word := range wordList {
		wordAnagrams := anagrams(word, dict)
		if len(wordAnagrams) > 1 {
			sort.Strings(wordAnagrams)
			result[strings.ToLower(strings.TrimSpace(word))] = wordAnagrams
		}
	}
	// [слово]["анаграмма_слова_1", ..., "анаграмма_слова_N"]
	return result
}
