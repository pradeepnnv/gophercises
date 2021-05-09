package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	inputFileName := "problems.csv"
	timeLimit := 30
	var err error
	if len(os.Args) > 2 {
		inputFileName = os.Args[1]
		timeLimit, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}

	inputFile, err := os.Open(inputFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	r := csv.NewReader(inputFile)
	count := 0
	corrAnsCount := 0
	for {
		t := time.NewTimer(time.Duration(timeLimit) * time.Second)
		defer t.Stop()
		go func() {
			<-t.C
			fmt.Printf("Sorry! Your time limit of %d exceeded!", timeLimit)
			os.Exit(0)
		}()

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		count++
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(record)
		fmt.Println("Your question is:", record[0])
		var ans string
		fmt.Scanln(&ans)

		if ans == record[1] {
			corrAnsCount++
		}

	}
	fmt.Printf("You answered %d questinos correctly out of %d", corrAnsCount, count)
}
