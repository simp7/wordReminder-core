package voca

type Word interface {
	Spell() Unit
	Mean() Unit
}
