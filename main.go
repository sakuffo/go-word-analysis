package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
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

func main() {
	// read the json file
	reviews := read_json_file("./data/Digital_Music_5.json")

	fmt.Println(reviews[0].ReviewText)
	fmt.Println("-----------")
	fmt.Println(reviews[1].ReviewText)
}
