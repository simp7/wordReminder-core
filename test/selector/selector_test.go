package selector

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/meaning"
	"github.com/simp7/wordReminder-core/voca/word"
	"github.com/simp7/wordReminder-core/voca/wordSet"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelector_Select(t *testing.T) {
	scenario := []struct {
		desc   string
		s      test.Selector
		amount int
		output voca.WordSet
	}{
		{"empty", &selector{0, wordSet.New("0")}, 0, wordSet.New("0")},
		{"empty but out of bound", &selector{0, wordSet.New("1")}, 5, wordSet.New("1")},
		{"a problem", &selector{1, wordSet.New("2", word.New("apple", meaning.Noun("사과")))}, 1, wordSet.New("2", word.New("apple", meaning.Noun("사과")))},
		{"problems", &selector{2, wordSet.New("3", word.New("apple", meaning.Noun("사과")), word.New("go", meaning.Verb("가다"), meaning.Noun("Go 언어")), word.New("bear", meaning.Noun("곰"), meaning.Verb("참다")))}, 3, wordSet.New("3", word.New("go", meaning.Verb("가다"), meaning.Noun("Go 언어")), word.New("bear", meaning.Noun("곰"), meaning.Verb("참다")), word.New("apple", meaning.Noun("사과")))},
		{"problems lower than original", &selector{4, wordSet.New("4", word.New("apple", meaning.Noun("사과")), word.New("go", meaning.Verb("가다"), meaning.Noun("Go 언어")), word.New("bear", meaning.Noun("곰"), meaning.Verb("참다")))}, 1, wordSet.New("4", word.New("go", meaning.Verb("가다"), meaning.Noun("Go 언어")))},
		{"problems more than original", &selector{4, wordSet.New("4", word.New("apple", meaning.Noun("사과")), word.New("go", meaning.Verb("가다"), meaning.Noun("Go 언어")), word.New("bear", meaning.Noun("곰"), meaning.Verb("참다")))}, 10, wordSet.New("4", word.New("go", meaning.Verb("가다"), meaning.Noun("Go 언어")), word.New("bear", meaning.Noun("곰"), meaning.Verb("참다")), word.New("apple", meaning.Noun("사과")))},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.s.Select(v.amount), v.desc)
	}

}
