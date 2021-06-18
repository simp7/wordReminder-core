package test

type Problems interface {
	Next() Problem
	HasNext() bool
	Size() int
}
