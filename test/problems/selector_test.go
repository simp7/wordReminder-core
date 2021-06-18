package problems

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/standard"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelector_Select(t *testing.T) {
	scenario := []struct {
		desc   string
		s      *selector
		amount int
		output voca.WordSet
	}{
		{"empty", &selector{0, standard.Words("0")}, 0, standard.Words("0")},
		{"empty but out of bound", &selector{0, standard.Words("1")}, 5, standard.Words("1")},
		{"a problem", &selector{1, standard.Words("2", standard.Word("apple", standard.Noun("사과")))}, 1, standard.Words("2", standard.Word("apple", standard.Noun("사과")))},
		{"problems", &selector{2, standard.Words("3", standard.Word("apple", standard.Noun("사과")), standard.Word("go", standard.Verb("가다"), standard.Noun("Go 언어")), standard.Word("bear", standard.Noun("곰"), standard.Verb("참다")))}, 3, standard.Words("3", standard.Word("go", standard.Verb("가다"), standard.Noun("Go 언어")), standard.Word("bear", standard.Noun("곰"), standard.Verb("참다")), standard.Word("apple", standard.Noun("사과")))},
		{"problems lower than original", &selector{4, standard.Words("4", standard.Word("apple", standard.Noun("사과")), standard.Word("go", standard.Verb("가다"), standard.Noun("Go 언어")), standard.Word("bear", standard.Noun("곰"), standard.Verb("참다")))}, 1, standard.Words("4", standard.Word("go", standard.Verb("가다"), standard.Noun("Go 언어")))},
		{"problems more than original", &selector{4, standard.Words("4", standard.Word("apple", standard.Noun("사과")), standard.Word("go", standard.Verb("가다"), standard.Noun("Go 언어")), standard.Word("bear", standard.Noun("곰"), standard.Verb("참다")))}, 10, standard.Words("4", standard.Word("go", standard.Verb("가다"), standard.Noun("Go 언어")), standard.Word("bear", standard.Noun("곰"), standard.Verb("참다")), standard.Word("apple", standard.Noun("사과")))},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.s.Select(v.amount), v.desc)
	}

}
