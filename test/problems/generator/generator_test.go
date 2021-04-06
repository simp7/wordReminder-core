package generator

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/test/problem"
	"github.com/simp7/wordReminder-core/test/problems"
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/meaning"
	"github.com/simp7/wordReminder-core/voca/word"
	"github.com/simp7/wordReminder-core/voca/wordSet"
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
		dictionary["apple"] = word.New("apple", meaning.Noun("사과"))
		dictionary["go"] = word.New("go", meaning.Noun("Go 언어"), meaning.Verb("가다"))
		dictionary["and"] = word.New("and", meaning.Conjunction("그리고"))
	})
	return dictionary[s]
}

func TestGenerator_Do(t *testing.T) {

	scenario := []struct {
		desc      string
		words     voca.WordSet
		generator problems.Generator
		output    test.Problems
	}{
		{"empty", wordSet.New("0"), Meaning(), nil},
		{"a word to test meaning", wordSet.New("1", wordOf("apple")), Meaning(), problems.Array(problem.Meaning(wordOf("apple")))},
		{"many words to test meanings", wordSet.New("2", wordOf("apple"), wordOf("go"), wordOf("and")), Meaning(), problems.Array(problem.Meaning(wordOf("apple")), problem.Meaning(wordOf("go")), problem.Meaning(wordOf("and")))},
		{"a word to test spelling", wordSet.New("3", wordOf("apple")), Spelling(), problems.Array(problem.Spelling(wordOf("apple")))},
		{"many words to test spellings", wordSet.New("4", wordOf("apple"), wordOf("go"), wordOf("and")), Spelling(), problems.Array(problem.Spelling(wordOf("apple")), problem.Spelling(wordOf("go")), problem.Spelling(wordOf("and")))},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.generator.Generate(v.words), v.desc)
	}

}
