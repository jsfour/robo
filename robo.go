package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/jsmootiv/robo/candidate"
	"github.com/jsmootiv/robo/question"
)

const EntryDelim = byte(',')

func processQuestions(myQuestions []*question.Question) []*question.Question {
	for _, q := range myQuestions {
		_, err := q.Answer(EntryDelim)
		if err != nil {
			fmt.Println(err)
			panic(1)
		}
	}
	return myQuestions
}

func writeAnswers(questions []*question.Question) {
	for _, q := range questions {
		fmt.Println(q.Prompt)
		fmt.Println(q.AnswerText)
		fmt.Printf("\n\n")
	}
}

func main() {
	data := strings.NewReader("temp")
	catColor := color.New(color.FgBlue)
	cand := candidate.BuildCandidate('\n')

	groupedQuestions, err := question.ReadQuestions(data)

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	for category, questions := range groupedQuestions {
		catColor.Println(category)
		processQuestions(questions)
	}

	cand.CollectNotes(EntryDelim)

	for category, questions := range groupedQuestions {
		color.Red("# %s Interview Notes", cand.Name)
		catColor.Println(category)
		writeAnswers(questions)
	}
}
