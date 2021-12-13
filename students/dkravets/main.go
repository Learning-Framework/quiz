package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fname := *flag.String("path", "problems.csv", "path to a csv file with quiz questions")
	timer := flag.Duration("timer", 30*time.Second, "Time given for a quiz game")
	flag.Parse()

	records, err := openCsv(fname)
	if err != nil {
		log.Fatalln("Openning csv file: ", err)
	}

	var correct, incorrect int
	done := make(chan string)
	fmt.Println("Please any key to start timer: ")
	fmt.Scanln()
	timer1 := time.NewTimer(*timer)
	go func() {
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
		done <- "done"
	}()
	<-timer1.C
	fmt.Printf("\nQuestions total: %d\t Correct answers: %d\n", len(records), correct)

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

// TODO: finish gracefully the programm if quiz was done before timer completed.
// TODO: implement timer for program execution.
