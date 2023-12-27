package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(unpackString("a4bc2d5e"))
}

func unpackString(input string) (output string) {
	left := 0
	right := 0
	arr := []rune(input)
	for right < len(arr) {
		if left == right {
			if right == len(arr)-1 {
				output += string(arr[left])
			}

			right++
			continue
		}
		if num, err := strconv.Atoi(string(arr[right])); err == nil {
			for j := 0; j < num; j++ {
				output += string(arr[left])
			}
			left = right + 1
			right++
		} else {
			output += string(arr[left])
			left = right
		}
	}
	return output
}
