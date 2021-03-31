package word

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/meaning"
	"github.com/simp7/wordReminder-core/voca/meanings"
	"github.com/simp7/wordReminder-core/voca/spelling"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWord_Mean(t *testing.T) {
	scenario := []struct {
		desc    string
		word    voca.Word
		meaning voca.Meanings
	}{
		{"single meaning", New("apple", meaning.Noun("사과")), meanings.New(meaning.Noun("사과"))},
		{"multi meanings", New("go", meaning.Verb("가다"), meaning.Noun("Go 언어")), meanings.New(meaning.Verb("가다"), meaning.Noun("Go 언어"))},
	}

	for _, v := range scenario {
		assert.Equal(t, v.meaning, v.word.Mean(), v.desc)
	}

}

func TestWord_Spell(t *testing.T) {
	scenario := []struct {
		desc    string
		word    voca.Word
		meaning voca.Spelling
	}{
		{"single meaning", New("apple", meaning.Noun("사과")), spelling.New("apple")},
		{"multi meanings", New("go", meaning.Verb("가다"), meaning.Noun("Go 언어")), spelling.New("go")},
	}

	for _, v := range scenario {
		assert.Equal(t, v.meaning, v.word.Spell(), v.desc)
	}

}
