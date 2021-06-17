package generator

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/test/problem"
	"github.com/simp7/wordReminder-core/test/problems"
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/standard"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var (
	once       = sync.Once{}
	dictionary map[string]voca.Word
)

func wordOf(s string) voca.Word {
	once.Do(func() {
		dictionary = make(map[string]voca.Word)
		dictionary["apple"] = standard.Word("apple", standard.Noun("사과"))
		dictionary["go"] = standard.Word("go", standard.Noun("Go 언어"), standard.Verb("가다"))
		dictionary["and"] = standard.Word("and", standard.Conjunction("그리고"))
	})
	return dictionary[s]
}

func TestGenerator_Do(t *testing.T) {

	scenario := []struct {
		desc      string
		words     voca.WordSet
		generator test.Generator
		output    test.Problems
	}{
		{"empty", standard.Words("0"), Meaning(), nil},
		{"a word to test meaning", standard.Words("1", wordOf("apple")), Meaning(), problems.Array(problem.Meaning(wordOf("apple")))},
		{"many words to test meanings", standard.Words("2", wordOf("apple"), wordOf("go"), wordOf("and")), Meaning(), problems.Array(problem.Meaning(wordOf("apple")), problem.Meaning(wordOf("go")), problem.Meaning(wordOf("and")))},
		{"a word to test spelling", standard.Words("3", wordOf("apple")), Spelling(), problems.Array(problem.Spelling(wordOf("apple")))},
		{"many words to test spellings", standard.Words("4", wordOf("apple"), wordOf("go"), wordOf("and")), Spelling(), problems.Array(problem.Spelling(wordOf("apple")), problem.Spelling(wordOf("go")), problem.Spelling(wordOf("and")))},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.generator.Generate(v.words), v.desc)
	}

}
