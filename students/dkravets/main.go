package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fname := *flag.String("path", "problems.csv", "path to a csv file with quiz questions")
	fl, err := os.Open(fname)
	if err != nil {
		log.Fatalln("Open quiz csv file: ", err)
	}

	csvr := csv.NewReader(fl)

	record, err := csvr.ReadAll()
	if err != nil {
		log.Fatalln("error reading csv : ", err)
	}
	for _, line := range record {
		fmt.Printf("Question: %s,\t answer: %s\n", line[0], line[1])
	}

}
