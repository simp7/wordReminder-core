package test

type Problems interface {
	Iterate(func(Problem))
	Add(Problem)
}
