package meanings

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/meaning"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMeanings_IsRight(t *testing.T) {

	scenario := []struct {
		desc     string
		meanings voca.Meanings
		ans      string
		output   bool
	}{
		{"apple", New(meaning.Noun("사과")), "사과", true},
		{"apple", New(meaning.Noun("사과")), "바나나", false},
		{"go", New(meaning.Verb("가다"), meaning.Noun("Go 언어")), "가다", true},
		{"go", New(meaning.Verb("가다"), meaning.Noun("Go 언어")), "Go 언어", true},
		{"trial and error", meaning.Idiom("시행착오"), "실패", false},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.meanings.IsRight(v.ans), v.desc)
	}

}

func TestMeanings_String(t *testing.T) {

	scenario := []struct {
		desc     string
		meanings voca.Meanings
		output   string
	}{
		{"apple", New(meaning.Noun("사과")), "사과"},
		{"go", New(meaning.Verb("가다"), meaning.Noun("Go 언어")), "가다, Go 언어"},
		{"beautiful", New(meaning.Adjective("아름다운")), "아름다운"},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.meanings.String(), v.desc)
	}

}
