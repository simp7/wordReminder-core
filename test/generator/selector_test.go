package generator

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca"
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
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.s.Select(v.amount), v.desc)
	}

}
