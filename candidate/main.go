package candidate

import (
	"fmt"

	"github.com/jsmootiv/robo/prompt"
)

type Candidate struct {
	Name  string
	Notes string
}

func (c *Candidate) CollectNotes(del byte) {
	notes, err := prompt.RunPrompt("General notes", del)
	if err != nil {
		fmt.Println("Error collecting notes on candidate")
	}

	c.Notes = notes
}

func BuildCandidate(del byte) *Candidate {
	candidate := Candidate{}
	name, err := prompt.RunPrompt("Candidate name", del)
	if err != nil {
		fmt.Println("Error reading candidate info")
	}

	candidate.Name = name
	return &candidate
}
