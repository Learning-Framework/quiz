package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fname := *flag.String("path", "problems.csv", "path to a csv file with quiz questions")
	flag.Parse()

	records, err := openCsv(fname)
	if err != nil {
		log.Fatalln("Openning csv file: ", err)
	}

	var correct, incorrect int
	for _, line := range records {
		fmt.Printf("Question: %s:\t", line[0])
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() != line[1] {
			fmt.Println("Incorrect! Correct answer is : ", line[1])
			incorrect += 1
		} else {
			correct += 1
			continue
		}
	}

	fmt.Printf("Questions total: %d\t Correct answers: %d\n", correct+incorrect, correct)

}

func openCsv(fname string) ([][]string, error) {
	var records [][]string
	fl, err := os.Open(fname)
	if err != nil {
		log.Fatalln("Open quiz csv file: ", err)
	}
	csvr := csv.NewReader(fl)
	records, err = csvr.ReadAll()
	return records, err
}

// TODO: part 2 of this quiz assignment has to be implemented
// TODO: Need to integrate goroutines to track time for each quiz question
