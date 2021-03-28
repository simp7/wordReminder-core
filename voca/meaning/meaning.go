package meaning

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/class"
)

type meaning struct {
	Class voca.Class
	Mean  string
}

func Noun(mean string) voca.Meaning {
	return meaning{Class: class.Noun, Mean: mean}
}

func Verb(mean string) voca.Meaning {
	return meaning{Class: class.Verb, Mean: mean}
}

func Adjective(mean string) voca.Meaning {
	return meaning{Class: class.Adjective, Mean: mean}
}

func Adverb(mean string) voca.Meaning {
	return meaning{Class: class.Adverb, Mean: mean}
}

func Conjunction(mean string) voca.Meaning {
	return meaning{Class: class.Conjunction, Mean: mean}
}

func Interjection(mean string) voca.Meaning {
	return meaning{Class: class.Interjection, Mean: mean}
}

func Preposition(mean string) voca.Meaning {
	return meaning{Class: class.Preposition, Mean: mean}
}

func Pronoun(mean string) voca.Meaning {
	return meaning{Class: class.Pronoun, Mean: mean}
}

func Idiom(mean string) voca.Meaning {
	return meaning{Class: class.Idiom, Mean: mean}
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
