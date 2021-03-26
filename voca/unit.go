package voca

//Unit is general interface for problem.
//Unit includes Meaning, Meanings, Spelling

type Unit interface {
	IsRight(string) bool
	String() string
}
