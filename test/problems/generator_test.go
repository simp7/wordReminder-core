package problems

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/standard"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerator_Generate(t *testing.T) {

	golang := standard.Word("go", standard.Verb("가다"), standard.Noun("Go 언어"))
	apple := standard.Word("apple", standard.Noun("사과"))

	scenario := []struct {
		input       voca.WordSet
		initializer func(voca.WordSet) *generator
		output      []test.Problem
		desc        string
	}{
		{standard.Words("0", golang), meaning, []test.Problem{test.Meaning(golang)}, "One meaning problem"},
		{standard.Words("0", golang, apple), meaning, []test.Problem{test.Meaning(golang), test.Meaning(apple)}, "Two meaning problems"},
		{standard.Words("0", golang), spelling, []test.Problem{test.Spelling(golang)}, "One meaning problem"},
		{standard.Words("0", golang, apple), spelling, []test.Problem{test.Spelling(golang), test.Spelling(apple)}, "Two meaning problems"},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.initializer(v.input).Generate(), v.desc)
	}

}
