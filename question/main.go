package question

import (
	"fmt"
	"io"

	"github.com/jsmootiv/robo/prompt"
)

type Question struct {
	Category   string
	Prompt     string
	AnswerText string
}

func (q *Question) Answer(delim byte) (*Question, error) {
	var err error
	fmt.Println("======================")
	fmt.Println("= " + q.Prompt)
	fmt.Println("======================")

	text, err := prompt.RunPrompt("== Notes", delim)
	fmt.Printf("\n\n")
	if err != nil {
		return q, err
	}
	q.AnswerText = text
	return q, err
}

func ReadQuestions(rawData io.Reader) (map[string][]*Question, error) {
	var err error
	out := make(map[string][]*Question)
	questions := []*Question{
		&Question{Category: "### foo", Prompt: "q1"},
		&Question{Category: "### foo", Prompt: "q2"},
		&Question{Category: "### foo", Prompt: "q3"},
	}
	readQuestions := len(questions)
	out["### foo"] = questions
	fmt.Println(fmt.Sprintf("Read %v questions\n", readQuestions))
	return out, err
}
