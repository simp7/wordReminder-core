package problems

import "github.com/simp7/wordReminder-core/test"

type array struct {
	data []test.Problem
	idx  int
}

func Array(pr ...test.Problem) *array {
	p := new(array)
	p.idx = -1
	p.data = pr
	return p
}

func (a *array) current() test.Problem {
	return a.data[a.idx]
}

func (a *array) next() bool {
	a.idx++
	return len(a.data) > a.idx
}

func (a *array) Iterate(f func(test.Problem)) {
	for a.next() {
		f(a.current())
	}
}

func (a *array) Add(problem test.Problem) {
	a.data = append(a.data, problem)
}
