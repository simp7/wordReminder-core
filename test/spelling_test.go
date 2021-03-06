package test

import (
	"github.com/simp7/wordReminder-core/voca/standard"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpellingProblem_IsCorrect(t *testing.T) {

	scenario := []struct {
		desc    string
		problem Problem
		input   string
		output  bool
	}{
		{"a meaning right", Spelling(standard.Word("apple", standard.Noun("사과"))), "apple", true},
		{"a meaning wrong", Spelling(standard.Word("apple", standard.Noun("사과"))), "pineapple", false},
		{"meanings right", Spelling(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다"))), "go", true},
		{"meanings wrong", Spelling(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다"))), "do", false},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.problem.IsCorrect(v.input), v.desc)
	}

}
