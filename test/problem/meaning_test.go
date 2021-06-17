package problem

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca/standard"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMeaningProblem_Question(t *testing.T) {

	scenario := []struct {
		desc    string
		problem test.Problem
		output  string
	}{
		{"a meaning", Meaning(standard.Word("apple", standard.Noun("사과"))), "apple"},
		{"meanings", Meaning(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다"))), "go"},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.problem.Question().String(), v.desc)
	}
}

func TestMeaningProblem_IsCorrect(t *testing.T) {

	scenario := []struct {
		desc    string
		problem test.Problem
		input   string
		output  bool
	}{
		{"a meaning right", Meaning(standard.Word("apple", standard.Noun("사과"))), "사과", true},
		{"a meaning wrong", Meaning(standard.Word("apple", standard.Noun("사과"))), "파인애플", false},
		{"meanings right 1", Meaning(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다"))), "가다", true},
		{"meanings right 2", Meaning(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다"))), "Go 언어", true},
		{"meanings wrong", Meaning(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다"))), "하다", false},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.problem.IsCorrect(v.input), v.desc)
	}

}
