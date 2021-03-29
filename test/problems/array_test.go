package problems

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/test/problem"
	"github.com/simp7/wordReminder-core/voca/meaning"
	"github.com/simp7/wordReminder-core/voca/word"
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
				ok = false
				return
			}
			if !problem.IsCorrect(v.input[i]) {
				ok = false
				return
			}
			i += 1
		})
		if ok {
			t.Logf("test \"%s\" has been passed!\n", v.desc)
		} else {
			t.Errorf("error in test \"%s\"\n ", v.desc)
		}
	}
}

func TestArray_Add(t *testing.T) {
	//scenario := []struct {
	//	desc     string
	//	problems test.Problems
	//	input    []string
	//}{
	//	{"single problem", Array(problem.Meaning(word.New("apple", meaning.Noun("사과")))), []string{"사과"}},
	//	{"multiple problems", Array(problem.Meaning(word.New("apple", meaning.Noun("사과"))), problem.Meaning(word.New("banana", meaning.Noun("바나나")))), []string{"사과", "바나나"}},
	//	{"multiple answers 1", Array(problem.Meaning(word.New("go", meaning.Noun("Go 언어"), meaning.Verb("가다")))), []string{"가다"}},
	//	{"multiple answers 2", Array(problem.Meaning(word.New("go", meaning.Noun("Go 언어"), meaning.Verb("가다")))), []string{"Go 언어"}},
	//	{"multiple problems and answers 1", Array(problem.Spelling(word.New("but", meaning.Conjunction("그러나"))), problem.Meaning(word.New("go", meaning.Noun("Go 언어"), meaning.Verb("가다")))), []string{"but","Go 언어"}},
	//	{"multiple problems and answers 2", Array(problem.Spelling(word.New("but", meaning.Conjunction("그러나"))), problem.Meaning(word.New("go", meaning.Noun("Go 언어"), meaning.Verb("가다")))), []string{"but", "가다"}},
	//}
}
