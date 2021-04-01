package meaning

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/class"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMeaning_IsRight(t *testing.T) {

	scenario := []struct {
		desc    string
		meaning voca.Meaning
		ans     string
		output  bool
	}{
		{"success1", Noun("사과"), "사과", true},
		{"success2", Verb("가다"), "가다", true},
		{"fail", Idiom("시행착오"), "실패", false},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.meaning.IsRight(v.ans), v.desc)
	}

}

func TestMeaning_String(t *testing.T) {

	scenario := []struct {
		desc    string
		meaning voca.Meaning
		output  string
	}{
		{"apple", Noun("사과"), "사과"},
		{"go", Verb("가다"), "가다"},
		{"beautiful", Adjective("아름다운"), "아름다운"},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.meaning.String())
	}

}

func TestMeaning_Type(t *testing.T) {

	scenario := []struct {
		desc    string
		meaning voca.Meaning
		output  voca.Class
	}{
		{"apple", Noun("apple"), class.Noun},
		{"he", Pronoun("그"), class.Pronoun},
		{"successfully", Adverb("성공적으로"), class.Adverb},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.meaning.Type(), v.desc)
	}

}
