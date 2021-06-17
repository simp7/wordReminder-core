package standard

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpelling_IsRight(t *testing.T) {

	scenario := []struct {
		desc     string
		spelling voca.Spelling
		ans      string
		output   bool
	}{
		{"right", Spelling("apple"), "apple", true},
		{"wrong", Spelling("apple"), "banana", false},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.spelling.IsRight(v.ans), v.desc)
	}

}

func TestSpelling_String(t *testing.T) {

	scenario := []struct {
		desc     string
		spelling voca.Spelling
		output   string
	}{
		{"spelling 1", Spelling("apple"), "apple"},
		{"spelling 2", Spelling("go"), "go"},
		{"spelling 3", Spelling("beautiful"), "beautiful"},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.spelling.String(), v.desc)
	}

}
