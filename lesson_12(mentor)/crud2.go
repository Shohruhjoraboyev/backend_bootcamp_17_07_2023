package main

import (
	"fmt"
	"sort"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Question struct {
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

func main() {
	questions := []Question{}
	questions = append(questions, Question{1, "1+2", []Variant{{1, "3", true}, {2, "4", false}, {3, "5", false}}})
	questions = append(questions, Question{2, "4-3", []Variant{{1, "3", false}, {2, "2", false}, {3, "1", true}}})
	questions = append(questions, Question{3, "7-3", []Variant{{1, "3", false}, {2, "4", true}, {3, "5", false}}})
	questions = append(questions, Question{4, "4+4", []Variant{{1, "8", true}, {2, "4", false}, {3, "5", false}}})
	questions = append(questions, Question{5, "9-3", []Variant{{1, "6", true}, {2, "4", false}, {3, "5", false}}})

	students := []Student{}
	students = append(students, Student{2, "Asadbek", []StudentAnswer{{1, "3"}, {2, "1"}, {3, "4"}, {4, "8"}, {5, "6"}}})
	students = append(students, Student{1, "Sardor", []StudentAnswer{{1, "3"}, {2, "2"}, {3, "4"}, {4, "8"}, {5, "6"}}})
	students = append(students, Student{3, "Odilbek", []StudentAnswer{{1, "1"}, {2, "1"}, {3, "1"}, {4, "2"}, {5, "2"}}})
	students = append(students, Student{4, "Valijon", []StudentAnswer{{1, "3"}, {2, "1"}, {3, "4"}, {4, "3"}, {5, "6"}}})
	students = append(students, Student{5, "Azamjon", []StudentAnswer{{1, "7"}, {2, "2"}, {3, "1"}, {4, "1"}, {5, "2"}}})
	students = append(students, Student{6, "Otabek", []StudentAnswer{{1, "3"}, {2, "2"}, {3, "4"}, {4, "8"}, {5, "5"}}})
	students = append(students, Student{7, "Oybek", []StudentAnswer{{1, "3"}, {2, "1"}, {3, "4"}, {4, "1"}, {5, "1"}}})

	correctAnswers := make(map[int]string)
	for _, question := range questions {
		for _, variant := range question.Options {
			if variant.Correct {
				correctAnswers[question.ID] = variant.Answer
			}
		}
	}

	studentsBall := make(map[string]int)
	for _, student := range students {
		for _, answer := range student.Answers {
			if correctAnswers[answer.QuestionID] == answer.Answer {
				studentsBall[student.Name] += 20
			} else {
				studentsBall[student.Name] += 0
			}
		}
	}

	stds := make([]string, 0, len(studentsBall))

	for student := range studentsBall {
		stds = append(stds, student)
	}

	sort.SliceStable(stds, func(i, j int) bool {
		return studentsBall[stds[i]] > studentsBall[stds[j]]
	})

	t := table.NewWriter()
	t.SetCaption("Natijalar")

	t.AppendHeader(table.Row{"O'RIN", "ISM", "BALL", "DARAJA"})
	for i, k := range stds {
		t.AppendRow(table.Row{i + 1, k, studentsBall[k], Ball(studentsBall[k])})
	}

	fmt.Println(t.Render())

}

func Ball(ball int) string {
	if ball < 55 {
		return "Failed"
	} else if ball < 71 {
		return "C"
	} else if ball < 86 {
		return "B"
	} else {
		return "A"
	}
}
