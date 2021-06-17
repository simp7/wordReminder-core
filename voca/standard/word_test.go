package standard

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWord_Mean(t *testing.T) {
	scenario := []struct {
		desc    string
		word    voca.Word
		meaning voca.Meanings
	}{
		{"single meaning", Word("apple", Noun("사과")), Meanings(Noun("사과"))},
		{"multi meanings", Word("go", Verb("가다"), Noun("Go 언어")), Meanings(Verb("가다"), Noun("Go 언어"))},
	}

	for _, v := range scenario {
		assert.Equal(t, v.meaning, v.word.Mean(), v.desc)
	}

}

func TestWord_Spell(t *testing.T) {
	scenario := []struct {
		desc     string
		word     voca.Word
		spelling voca.Spelling
	}{
		{"spelling 1", Word("apple", Noun("사과")), Spelling("apple")},
		{"spelling 2", Word("go", Verb("가다"), Noun("Go 언어")), Spelling("go")},
	}

	for _, v := range scenario {
		assert.Equal(t, v.spelling, v.word.Spell(), v.desc)
	}

}
