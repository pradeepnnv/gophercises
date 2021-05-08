package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	inputFileName := "problems.csv"
	if len(os.Args) > 1 {
		inputFileName = os.Args[1]
	} else {
		fmt.Println(`Input file name is not provided. Using default file name of "problems.csv".`)
	}

	inputFile, err := os.Open(inputFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	r := csv.NewReader(inputFile)
	count := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(record)
		fmt.Println("Your question is :", record[0])
		var ans string
		fmt.Scanln(&ans)

		if ans == record[1] {
			count++
		}
	}
	fmt.Printf("Number of correct answers is %d", count)
}
