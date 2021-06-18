package problems

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
)

type generator struct {
	words voca.WordSet
}

func meaning(words voca.WordSet) *generator {

	g := new(generator)
	g.words = words

	return g

}

func (g *generator) Generate() {

	problems := make([]test.Problem, g.words.Len())

	for i := 0; i < g.words.Len(); i++ {
		problems[i] = test.Meaning(g.words.Get(i))
	}

}
