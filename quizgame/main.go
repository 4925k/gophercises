package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	//flags initialization
	path := flag.String("path", "problems.csv", "path to your questions")
	timeLimit := flag.Int("timeLimit", 60, "timer for quiz")
	flag.Parse()

	//start the quiz
	startQuiz(*path, *timeLimit)
}

func startQuiz(path string, timeLimit int) {
	//get questions from the csv file
	qa := fetchQuestions(path)

	//start a timer
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	correct := 0 //store correct answers

questionLoop:
	for n, line := range qa {
		//ask question
		fmt.Printf("Question %v: %v \nAnswer: ", n+1, line[0])

		//ask for answer
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer
		}()

		//select case for if time is over before the question is answered
		select {
		case <-timer.C:
			fmt.Println("\nTime is over.")
			break questionLoop
		case answer := <-answerCh:
			if answer == line[1] {
				correct++
			}
		}
	}

	//print results
	fmt.Printf("Correct Answers: %v || Total Questions: %v", correct, len(qa))
}

//fetchQuestions parses the csv file
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

//quit exits from the program after showing given message
func quit(msg error) {
	fmt.Printf("ERROR %v", msg)
	os.Exit(1)
}
