package main

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

}
