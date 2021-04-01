package spelling

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
		{"apple", New("apple"), "apple", true},
		{"apple", New("apple"), "banana", false},
		{"go", New("go"), "go", true},
		{"trial and error", New("trial and error"), "trials", false},
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
		{"apple", New("apple"), "apple"},
		{"go", New("go"), "go"},
		{"beautiful", New("beautiful"), "beautiful"},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.spelling.String(), v.desc)
	}

}
