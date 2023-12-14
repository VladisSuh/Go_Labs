package main

import (
	"fmt"
	"strings"
)

func isAnagram(a, b string) bool {
	aFreq := make(map[rune]int)
	for _, c := range a {
		aFreq[c]++
	}
	bFreq := make(map[rune]int)
	for _, c := range b {
		bFreq[c]++
	}
	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}
	for k, v := range bFreq {
		if aFreq[k] != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Введите первую строку: ")
	var firstString string
	fmt.Scanln(&firstString)
	firstString = strings.TrimSpace(firstString)

	fmt.Print("Введите вторую строку: ")
	var secondString string
	fmt.Scanln(&secondString)
	secondString = strings.TrimSpace(secondString)

	if isAnagram(firstString, secondString) {
		fmt.Println("Эти строки являются анаграммами.")
	} else {
		fmt.Println("Эти строки не являются анаграммами.")
	}
}
