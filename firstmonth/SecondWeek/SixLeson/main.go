package main

import (
	"fmt"
	c "Crud"
)

type Questions struct {
	Id      int
	Text    string
	Options []Variant
}

type Variant struct {
	Id      int
	Answer  string
	Correct bool
}

type Student struct {
	Id      int
	Name    string
	Answers []Student_Answer
}

type Student_Answer struct {
	QuestionId int
	Answer     string
}

func main() {
	var variant = []Variant{
		{
			Id:      4,
			Answer:  "Toshkent",
			Correct: true,
		},
		{
			Id:      4,
			Answer:  "Andijon",
			Correct: false,
		},
		{
			Id:      4,
			Answer:  "Samarqand",
			Correct: false,
		},
		{
			Id:      4,
			Answer:  "Termiz",
			Correct: false,
		},
		//
		{
			Id:      1,
			Answer:  "Toshkent",
			Correct: true,
		},
		{
			Id:      1,
			Answer:  "Andijon",
			Correct: false,
		},
		{
			Id:      1,
			Answer:  "Samarqand",
			Correct: false,
		},
		{
			Id:      1,
			Answer:  "Termiz",
			Correct: false,
		},
		//
		{
			Id:      2,
			Answer:  "Toshkent",
			Correct: true,
		},
		{
			Id:      2,
			Answer:  "Andijon",
			Correct: false,
		},
		{
			Id:      2,
			Answer:  "Samarqand",
			Correct: false,
		},
		{
			Id:      2,
			Answer:  "Termiz",
			Correct: false,
		},
		//
		{
			Id:      3,
			Answer:  "Toshkent",
			Correct: true,
		},
		{
			Id:      3,
			Answer:  "Andijon",
			Correct: false,
		},
		{
			Id:      3,
			Answer:  "Samarqand",
			Correct: false,
		},
		{
			Id:      3,
			Answer:  "Termiz",
			Correct: false,
		},
		//
	}

	var questions = []Questions{
		{
			Id:      1,
			Text:    "O'zbekiston respublikasining poytaxti qaysi viloyat ? ",
			Options: variant,
		},
		{
			Id:      2,
			Text:    "O'zbekiston respublikasining poytaxti qaysi viloyat ? ",
			Options: variant,
		},
		{
			Id:      3,
			Text:    "O'zbekiston respublikasining poytaxti qaysi viloyat ? ",
			Options: variant,
		},
		{
			Id:      4,
			Text:    "O'zbekiston respublikasining poytaxti qaysi viloyat ? ",
			Options: variant,
		},
	}

	var student_Answer = []Student_Answer{
		{
			QuestionId: 1,
			Answer:     "Andijan",
		},
		{
			QuestionId: 2,
			Answer:     "Toshkent",
		},
		{
			QuestionId: 3,
			Answer:     "Toshkent",
		},
		{
			QuestionId: 4,
			Answer:     "Toshkent",
		},
	}

	var student = []Student{
		{
			Id:      1,
			Name:    "Sarvarbek",
			Answers: student_Answer,
		},
	}

	grade := calculateGrade(questions, student, student[0].Answers)
	fmt.Println("Grade:", grade)

}
func calculateGrade(questions []Questions, students []Student, studentAnswers []Student_Answer) string {

	caMap := make(map[int]string)
	for _, q := range questions {
		for _, opt := range q.Options {
			if opt.Correct {
				caMap[q.Id] = opt.Answer
			}
		}
	}

	resMap := make(map[int]int)
	for _, answer := range studentAnswers {
		cA, _ := caMap[answer.QuestionId]
		if answer.Answer == cA {
			for _, v := range students {
				resMap[v.Id]++

			}
			// resMap[students[0].Id]++
		}
	}

	questCount := len(questions)
	result := (float32(resMap[students[0].Id]) / float32(questCount)) * 100

	if result >= 86 && result <= 100 {
		return "A"
	} else if result >= 71 && result < 85 {
		return "B"
	} else if result >= 55 && result <= 70 {
		return "C"
	} else {
		return "Failed"
	}
}

fmt.Println(Crud.Crud)