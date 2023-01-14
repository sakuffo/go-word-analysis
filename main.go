package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Review struct {
	ReviewerID   string  `json:"reviewerID"`
	Asin         string  `json:"asin"`
	ReviewerName string  `json:"reviewerName"`
	Helpful      [2]int  `json:"helpful"`
	ReviewText   string  `json:"reviewText"`
	Overall      float32 `json:"overall"`
	Summary      string  `json:"summary"`
	UnixReviewTm int     `json:"unixReviewTime"`
	ReviewTime   string  `json:"reviewTime"`
	Tokens       []string
	WordCount    map[string]int
}

type Dataset struct {
	filepath string
	reviews  []Review
}

func (dataset *Dataset) read_json_file() {
	content, err := os.Open(dataset.filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var review Review
		err := json.Unmarshal([]byte(scanner.Text()), &review)
		if err != nil {
			log.Fatal(err)
		}

		dataset.reviews = append(dataset.reviews, review)
	}
}

func read_json_file(filepath string) []Review {
	// read the json file using the os package
	content, err := os.Open(filepath)
	// if we have an error, we log the error and exit the program
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file until the read_json_file function finishes
	defer content.Close()
	// create a scanner variable that we will use to iterate through the reviews
	scanner := bufio.NewScanner(content)
	// split the content of the file based on lines (each line is a review)
	scanner.Split(bufio.ScanLines)

	var reviews []Review
	for scanner.Scan() {
		var review Review
		err := json.Unmarshal([]byte(scanner.Text()), &review)
		if err != nil {
			log.Fatal(err)
		}
		reviews = append(reviews, review)
	}

	return reviews
}

func (dataset *Dataset) tokenize() {
	for i := range dataset.reviews {
		dataset.reviews[i].Tokens = tokenize(dataset.reviews[i].ReviewText)
	}
}

func tokenize(text string) []string {

	re := regexp.MustCompile(`[[:punct:]]`)

	w := re.ReplaceAllString(text, " ")

	w = strings.ToLower(w)

	tokens := strings.Fields(w)

	return tokens
}

func count_words(words []string) map[string]int {

	word_count := make(map[string]int)

	for i := range words {
		if _, ok := word_count[words[i]]; ok {
			word_count[words[i]] += 1
		} else {
			word_count[words[i]] = 1
		}
	}
	return word_count
}

func (dataset *Dataset) count_words() {
	for i := range dataset.reviews {
		dataset.reviews[i].WordCount = count_words(dataset.reviews[i].Tokens)
	}
}

func main() {
	dataset := Dataset{filepath: "./data/Digital_Music_5.json"}
	dataset.read_json_file()
	dataset.tokenize()
	dataset.count_words()
	for i := range dataset.reviews {
		fmt.Println(dataset.reviews[i].ReviewText)
		fmt.Println("---")
		fmt.Println(dataset.reviews[i].Tokens)
		fmt.Println("---")
		fmt.Println(dataset.reviews[i].WordCount)
		fmt.Println("=================================")

		if i >= 2 {
			break
		}
	}
}
