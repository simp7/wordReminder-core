package test

type Creator interface {
	Create(amount int) Problems
	CreateAll() Problems
}
