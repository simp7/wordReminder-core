package test

type Exam interface {
	Iterate(func(Problem))
	Amount() int
	Right() int
}
