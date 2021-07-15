package test

type Exam struct {
	Problems
	Total   int
	Solved  int
	Correct int
}

func New(problems Problems) *Exam {

	e := new(Exam)
	e.Problems = problems

	e.Total = problems.Size()

	return e

}

func (e *Exam) Iterate(process func(problem Problem) bool) {
	for e.HasNext() {
		problem := e.Next()
		isRight := process(problem)
		if isRight {
			e.Correct++
		}
		e.Solved++
	}
}
