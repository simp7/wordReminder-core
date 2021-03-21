package test

type Problem interface {
	Question() string
	IsCorrect(userAnswer string) bool
}


