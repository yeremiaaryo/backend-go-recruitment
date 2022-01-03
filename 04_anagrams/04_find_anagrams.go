package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	result := findAnagrams(input)

	for _, v := range result {
		fmt.Println(v)
	}
}

func findAnagrams(input []string) map[string][]string {
	mapInput := make(map[string][]string, 0)
	for _, v := range input {
		word := strings.Split(v, "")
		sort.Strings(word)
		sortedWords := strings.Join(word, "")
		mapInput[sortedWords] = append(mapInput[sortedWords], v)
	}

	return mapInput
}
