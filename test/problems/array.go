package problems

import "wordReminder-core/test"

type array struct {
	data []test.Problem
	idx  int
}

func Array(pr... test.Problem) test.Problems {
	p := new(array)
	p.idx = 0
	p.data = pr
	return p
}

func (a *array) Current() test.Problem {
	return a.data[a.idx]
}

func (a *array) Next() {
	a.idx ++
}