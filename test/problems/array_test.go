package problems

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/test/problem"
	"github.com/simp7/wordReminder-core/voca/meaning"
	"github.com/simp7/wordReminder-core/voca/word"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray_Iterate(t *testing.T) {

	scenario := []struct {
		desc     string
		problems test.Problems
		input    []string
	}{
		{"single problem", Array(problem.Meaning(word.New("apple", meaning.Noun("사과")))), []string{"사과"}},
		{"multiple problems", Array(problem.Meaning(word.New("apple", meaning.Noun("사과"))), problem.Meaning(word.New("banana", meaning.Noun("바나나")))), []string{"사과", "바나나"}},
		{"multiple answers 1", Array(problem.Meaning(word.New("go", meaning.Noun("Go 언어"), meaning.Verb("가다")))), []string{"가다"}},
		{"multiple answers 2", Array(problem.Meaning(word.New("go", meaning.Noun("Go 언어"), meaning.Verb("가다")))), []string{"Go 언어"}},
		{"multiple problems and answers 1", Array(problem.Spelling(word.New("but", meaning.Conjunction("그러나"))), problem.Meaning(word.New("go", meaning.Noun("Go 언어"), meaning.Verb("가다")))), []string{"but", "Go 언어"}},
		{"multiple problems and answers 2", Array(problem.Spelling(word.New("but", meaning.Conjunction("그러나"))), problem.Meaning(word.New("go", meaning.Noun("Go 언어"), meaning.Verb("가다")))), []string{"but", "가다"}},
	}

	for _, v := range scenario {
		i := 0
		ok := true
		v.problems.Iterate(func(problem test.Problem) {
			if len(v.input) == i {
				t.Errorf("error in test \"%s\": index out of range ", v.desc)
				ok = false
				return
			}
			if !problem.IsCorrect(v.input[i]) {
				t.Errorf("error in test \"%s\": problem.Iscorrect(%s) returns false\n ", v.desc, v.input[i])
				ok = false
				return
			}
			i += 1
		})
		if ok {
			t.Logf("test %d has been passed!", i)
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
		{"single problem", Array(problem.Meaning(word.New("apple", meaning.Noun("사과")))), problem.Meaning(word.New("zebra", meaning.Noun("얼룩말"))), Array(problem.Meaning(word.New("apple", meaning.Noun("사과"))), problem.Meaning(word.New("zebra", meaning.Noun("얼룩말"))))},
		{"multiple problems", Array(problem.Meaning(word.New("apple", meaning.Noun("사과"))), problem.Spelling(word.New("in", meaning.Preposition("~에")))), problem.Meaning(word.New("wow", meaning.Interjection("놀람"))), Array(problem.Meaning(word.New("apple", meaning.Noun("사과"))), problem.Spelling(word.New("in", meaning.Preposition("~에"))), problem.Meaning(word.New("wow", meaning.Interjection("놀람"))))},
	}

	for _, v := range scenario {
		v.problems.Add(v.input)
		assert.Equal(t, v.output, v.problems, v.desc)
	}

}
