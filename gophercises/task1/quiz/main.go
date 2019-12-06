package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "../problems.csv", "a file with questions")
	timeLimit := flag.Int("limit", 15, "the time limit for quiz")
	randomProblems := flag.Bool("random", true, "mix promblems order")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	defer file.Close()
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Fail to parse the CSV file.")
	}
	if len(lines) == 0 {
		exit("No problems in the CSV file.")
	}
	problems := parseLines(lines)
	if *randomProblems {
		problems = shuffle(problems)
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYour scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerChan:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("Your scored %d out of %d.\n", correct, len(problems))
}

func shuffle(p []problem) []problem {
	rand.Seed(time.Now().UnixNano())
	for i := range p {
		r := rand.Intn(len(p))
		p[i], p[r] = p[r], p[i]
	}
	return p
}

func parseLines(lines [][]string) []problem {
	res := make([]problem, len(lines))
	for i, line := range lines {
		res[i] = problem{
			q: line[0],
			a: strings.ToLower(strings.TrimSpace(line[1])),
		}
	}
	return res
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
