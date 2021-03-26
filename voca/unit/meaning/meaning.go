package meaning

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/class"
	"github.com/simp7/wordReminder-core/voca/unit"
)

type meaning struct {
	Class voca.Class
	Mean  string
}

func Noun(mean string) unit.Meaning {
	return meaning{Class: class.Noun, Mean: mean}
}

func Verb(mean string) unit.Meaning {
	return meaning{Class: class.Verb, Mean: mean}
}

func Adjective(mean string) unit.Meaning {
	return meaning{Class: class.Adjective, Mean: mean}
}

func AdVerb(mean string) unit.Meaning {
	return meaning{Class: class.Adverb, Mean: mean}
}

func Conjunction(mean string) unit.Meaning {
	return meaning{Class: class.Conjunction, Mean: mean}
}

func Interjection(mean string) unit.Meaning {
	return meaning{Class: class.Interjection, Mean: mean}
}

func Preposition(mean string) unit.Meaning {
	return meaning{Class: class.Preposition, Mean: mean}
}

func Pronoun(mean string) unit.Meaning {
	return meaning{Class: class.Pronoun, Mean: mean}
}

func Idiom(mean string) unit.Meaning {
	return meaning{Class: class.Idiom, Mean: mean}
}

func (m meaning) IsRight(input string) bool {
	return m.Mean == input
}

func (m meaning) Format() string {
	return m.Mean
}
