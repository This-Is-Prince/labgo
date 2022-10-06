package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func problemPuller(fileName string) ([]Problem, error) {
	if file, err := os.Open(fileName); err == nil {
		csvReader := csv.NewReader(file)
		if lines, err := csvReader.ReadAll(); err == nil {
			return parseProblem(lines), nil
		} else {
			return nil, fmt.Errorf("error in reading data in csv format from %s file; %s", fileName, err.Error())
		}
	} else {
		return nil, fmt.Errorf("error in opening %s file; %s", fileName, err.Error())
	}
}

func parseProblem(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for i, line := range lines {
		problems[i] = Problem{q: line[0], a: line[1]}
	}
	return problems
}

type Problem struct {
	q string
	a string
}

func main() {
	fmt.Println("Go Quiz App")

	fName := flag.String("f", "quiz.csv", "path of csv file")
	timer := flag.Int("t", 30, "timer for the quiz")
	flag.Parse()

	problems, err := problemPuller(*fName)
	if err != nil {
		log.Fatalf("something went wrong:%s\n", err.Error())
	}

	correctAns := 0

	tmr := time.NewTimer(time.Duration(*timer) * time.Second)
	answerChannel := make(chan string)

problemLoop:
	for i, pbm := range problems {
		var ans string
		fmt.Printf("Problem %d: %s=", i+1, pbm.q)
		go func() {
			fmt.Scanf("%s", &ans)
			answerChannel <- ans
		}()
		select {
		case <-tmr.C:
			fmt.Println()
			break problemLoop
		case ans := <-answerChannel:
			if ans == pbm.a {
				correctAns++
			}
			if i == len(problems)-1 {
				close(answerChannel)
			}
		}
	}

	fmt.Printf("Your result is %d out of %d\n", correctAns, len(problems))
	fmt.Println("Press enter to exit")
	<-answerChannel
}
