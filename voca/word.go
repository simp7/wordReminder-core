package voca

type Word interface {
	Spell() Spelling
	Mean() Meanings
}
