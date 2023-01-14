package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := "Hello, Sean! -How are you?"
	fmt.Println("original string: " + text)
	result := removePunctuation(text)
	fmt.Println("string after removing punctuation: " + result)
	tokens := strings.Fields(result)

	fmt.Print("Tokens: ")
	fmt.Println(tokens)
}

func removePunctuation(s string) string {
	re := regexp.MustCompile(`[[:punct:]]`)
	w := re.ReplaceAllString(s, " ")
	return strings.ToLower(w)
}
