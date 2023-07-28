package main

import "fmt"

type Questions struct {
	ID      int
	Text    string
	Options []Variant
}

type Variant struct {
	ID      int
	Answer  string
	Correct bool
}

type Student struct {
	ID      int
	Name    string
	Answers []StudentAnswer
}

type StudentAnswer struct {
	QuestionID int
	Answer     string
}

type Results struct {
	StudentId      int
	CorrectAnswers int
}

func main() {
	// Create some questions
	q1 := Questions{
		ID:   1,
		Text: "What is the capital of France?",
		Options: []Variant{
			{ID: 1, Answer: "London", Correct: false},
			{ID: 2, Answer: "Paris", Correct: true},
			{ID: 3, Answer: "Berlin", Correct: false},
		},
	}
	q2 := Questions{
		ID:   2,
		Text: "What is the largest planet in the solar system?",
		Options: []Variant{
			{ID: 1, Answer: "Mars", Correct: false},
			{ID: 2, Answer: "Jupiter", Correct: true},
			{ID: 3, Answer: "Venus", Correct: false},
		},
	}

	// Create some students and their answers
	s1 := Student{
		ID:   1,
		Name: "Alice",
		Answers: []StudentAnswer{
			{QuestionID: 1, Answer: "Paris"},
			{QuestionID: 2, Answer: "Mars"},
		},
	}
	s2 := Student{
		ID:   2,
		Name: "Bob",
		Answers: []StudentAnswer{
			{QuestionID: 1, Answer: "London"},
			{QuestionID: 2, Answer: "Jupiter"},
		},
	}

	// Store the questions and students in a map
	data := make(map[int]interface{})
	data[q1.ID] = q1
	data[q2.ID] = q2
	data[s1.ID] = s1
	data[s2.ID] = s2

	// Calculate and print the number of correct answers found by each student
	results := make(map[int]int)
	for _, v := range data {
		switch t := v.(type) {
		case Student:
			for _, a := range t.Answers {
				q, ok := data[a.QuestionID].(Questions)
				if !ok {
					continue
				}
				if findCorrectAnswer(q.Options) == a.Answer {
					results[t.ID]++
				}
			}
		}
	}
	fmt.Println("Results:")
	for k, v := range results {
		fmt.Printf("Student %d found %d correct answers\n", k, v)
	}
}

// Helper function to find the correct answer to a question
func findCorrectAnswer(options []Variant) string {
	for _, o := range options {
		if o.Correct {
			return o.Answer
		}
	}
	return ""
}
