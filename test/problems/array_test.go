package problems

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/test/problem"
	"github.com/simp7/wordReminder-core/voca/standard"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray_Iterate(t *testing.T) {

	scenario := []struct {
		desc     string
		problems *array
		input    []string
	}{
		{"single problem", Array(problem.Meaning(standard.Word("apple", standard.Noun("사과")))), []string{"사과"}},
		{"multiple problems", Array(problem.Meaning(standard.Word("apple", standard.Noun("사과"))), problem.Meaning(standard.Word("banana", standard.Noun("바나나")))), []string{"사과", "바나나"}},
		{"multiple answers 1", Array(problem.Meaning(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다")))), []string{"가다"}},
		{"multiple answers 2", Array(problem.Meaning(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다")))), []string{"Go 언어"}},
		{"multiple problems and answers 1", Array(problem.Spelling(standard.Word("but", standard.Conjunction("그러나"))), problem.Meaning(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다")))), []string{"but", "Go 언어"}},
		{"multiple problems and answers 2", Array(problem.Spelling(standard.Word("but", standard.Conjunction("그러나"))), problem.Meaning(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다")))), []string{"but", "가다"}},
	}

	for _, v := range scenario {
		assert.Len(t, v.problems.data, len(v.input), v.desc)
		for i := 0; v.problems.HasNext(); i++ {
			assert.True(t, v.problems.Next().IsCorrect(v.input[i]), v.desc)
		}
	}

}

func TestArray_Add(t *testing.T) {

	scenario := []struct {
		desc     string
		problems test.Problems
		input    test.Problem
		output   test.Problems
	}{
		{"single problem", Array(problem.Meaning(standard.Word("apple", standard.Noun("사과")))), problem.Meaning(standard.Word("zebra", standard.Noun("얼룩말"))), Array(problem.Meaning(standard.Word("apple", standard.Noun("사과"))), problem.Meaning(standard.Word("zebra", standard.Noun("얼룩말"))))},
		{"multiple problems", Array(problem.Meaning(standard.Word("apple", standard.Noun("사과"))), problem.Spelling(standard.Word("in", standard.Preposition("~에")))), problem.Meaning(standard.Word("wow", standard.Interjection("놀람"))), Array(problem.Meaning(standard.Word("apple", standard.Noun("사과"))), problem.Spelling(standard.Word("in", standard.Preposition("~에"))), problem.Meaning(standard.Word("wow", standard.Interjection("놀람"))))},
	}

	for _, v := range scenario {
		v.problems.Add(v.input)
		assert.Equal(t, v.output, v.problems, v.desc)
	}

}
