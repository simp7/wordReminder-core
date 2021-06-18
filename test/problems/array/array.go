package array

import "github.com/simp7/wordReminder-core/test"

type array struct {
	data []test.Problem
	idx  int
}

func Array(pr ...test.Problem) test.Problems {
	p := new(array)
	p.idx = -1
	p.data = pr
	return p
}

func (a *array) Next() test.Problem {
	a.idx++
	return a.data[a.idx]
}

func (a *array) HasNext() bool {
	return len(a.data)-1 != a.idx
}

func (a *array) Size() int {
	return len(a.data)
}
