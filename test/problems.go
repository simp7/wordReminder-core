package test

type Problems interface {
	Next()
	Current() Problem
}