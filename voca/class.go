package voca

//Class represents class of word like Verb.
type Class uint8

const (
	Noun Class = iota
	Verb
	Adverb
	Adjective
	Conjunction
	Pronoun
	Preposition
	Interjection
	Idiom
)
