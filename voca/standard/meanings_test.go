package standard

import (
	"github.com/simp7/wordReminder-core/voca"
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
		{"a meaning right", Meanings(Noun("사과")), "사과", true},
		{"a meaning wrong", Meanings(Noun("사과")), "바나나", false},
		{"meanings right 1", Meanings(Verb("가다"), Noun("Go 언어")), "가다", true},
		{"meanings right 2", Meanings(Verb("가다"), Noun("Go 언어")), "Go 언어", true},
		{"meanings wrong", Meanings(Verb("가다"), Noun("Go 언어")), "하다", false},
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
		{"a meaning", Meanings(Noun("사과")), "사과"},
		{"meanings", Meanings(Verb("가다"), Noun("Go 언어")), "가다, Go 언어"},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.meanings.String(), v.desc)
	}

}
