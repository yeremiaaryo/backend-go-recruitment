package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "hello (world)"
	fmt.Println("Result:", findFirstStringInBracket(input))
}

func findFirstStringInBracket(str string) string {
	indexFirstBracketFound := strings.Index(str, "(")
	indexClosingBracketFound := strings.Index(str, ")")
	if indexFirstBracketFound >= 0 && indexClosingBracketFound >= 0 {
		return str[indexFirstBracketFound+1 : indexClosingBracketFound]
	}
	return ""
}
