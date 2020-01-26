package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{q: line[0], a: line[1]}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "csv file for problems")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Printf("!Failed to parse %s", *csvFile)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failed to print")
	}
	correct := 0
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for i, p := range problems {
		fmt.Printf("Problem no #%d : %s = \n", i+1, p.q)
		ansChan := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansChan <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d \n", correct, len(problems))
			return
		case answer := <-ansChan:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("\nYou scored %d out of %d \n", correct, len(problems))
}
