package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

var path *string
var time *int

func init() {
	//flags initialization
	path = flag.String("path", "problems.csv", "path to your questions")
	time = flag.Int("time", 60, "timer for quiz")
	flag.Parse()

}

func main() {
	//start the quiz
	startQuiz(*path, *time)
}

func startQuiz(path string, time int) {
	qa := fetchQuestions(path)

	//loop and ask all the questions
	correct := 0 //store the number of correct answer
	var answer string
	for n, line := range qa {
		fmt.Printf("Question %v: %v \nAnswer: ", n+1, line[0])
		fmt.Scanln(&answer)
		if answer == line[1] {
			correct++
		}
	}
	fmt.Printf("Correct Answers: %v || Total Questions: %v", correct, len(qa))
}

func fetchQuestions(path string) [][]string {
	//open problems file for questions and answers
	csvFile, err := os.Open(path)
	if err != nil {
		quit(err)
	}

	//read all the questions and answers
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		quit(err)
	}

	return csvLines
}

func quit(msg error) {
	fmt.Printf("ERROR %v", msg)
	os.Exit(1)
}
