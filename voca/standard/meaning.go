package standard

import (
	"github.com/simp7/wordReminder-core/voca"
)

type meaning struct {
	Class voca.Class
	Mean  string
}

func Noun(mean string) voca.Meaning {
	return meaning{Class: voca.Noun, Mean: mean}
}

func Verb(mean string) voca.Meaning {
	return meaning{Class: voca.Verb, Mean: mean}
}

func Adjective(mean string) voca.Meaning {
	return meaning{Class: voca.Adjective, Mean: mean}
}

func Adverb(mean string) voca.Meaning {
	return meaning{Class: voca.Adverb, Mean: mean}
}

func Conjunction(mean string) voca.Meaning {
	return meaning{Class: voca.Conjunction, Mean: mean}
}

func Interjection(mean string) voca.Meaning {
	return meaning{Class: voca.Interjection, Mean: mean}
}

func Preposition(mean string) voca.Meaning {
	return meaning{Class: voca.Preposition, Mean: mean}
}

func Pronoun(mean string) voca.Meaning {
	return meaning{Class: voca.Pronoun, Mean: mean}
}

func Idiom(mean string) voca.Meaning {
	return meaning{Class: voca.Idiom, Mean: mean}
}

func (m meaning) IsRight(input string) bool {
	return m.Mean == input
}

func (m meaning) String() string {
	return m.Mean
}

func (m meaning) Type() voca.Class {
	return m.Class
}
