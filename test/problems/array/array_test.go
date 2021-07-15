package array

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca/standard"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray_Iterate(t *testing.T) {

	scenario := []struct {
		desc     string
		problems test.Problems
		input    []string
	}{
		{"single problem", Array(test.Meaning(standard.Word("apple", standard.Noun("사과")))), []string{"사과"}},
		{"multiple problems", Array(test.Meaning(standard.Word("apple", standard.Noun("사과"))), test.Meaning(standard.Word("banana", standard.Noun("바나나")))), []string{"사과", "바나나"}},
		{"multiple answers 1", Array(test.Meaning(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다")))), []string{"가다"}},
		{"multiple answers 2", Array(test.Meaning(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다")))), []string{"Go 언어"}},
		{"multiple problems and answers 1", Array(test.Spelling(standard.Word("but", standard.Conjunction("그러나"))), test.Meaning(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다")))), []string{"but", "Go 언어"}},
		{"multiple problems and answers 2", Array(test.Spelling(standard.Word("but", standard.Conjunction("그러나"))), test.Meaning(standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다")))), []string{"but", "가다"}},
	}

	for _, v := range scenario {
		assert.Equal(t, v.problems.Size(), len(v.input), v.desc)
		for i := 0; v.problems.HasNext(); i++ {
			assert.True(t, v.problems.Next().IsCorrect(v.input[i]), v.desc)
		}
	}

}
