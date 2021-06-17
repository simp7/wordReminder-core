package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStreak_Increment(t *testing.T) {

	scenario := []struct {
		desc   string
		input  uint
		output uint
	}{
		{"zero++", 0, 1},
		{"one++", 1, 2},
		{"large++", 2021, 2022},
	}

	for _, v := range scenario {
		input := Streak(v.input)
		input.Increment()
		output := Streak(v.output)

		assert.Equal(t, output, input, v.desc)
	}

}

func TestStreak_Decrement(t *testing.T) {
	scenario := []struct {
		desc   string
		input  uint
		output uint
	}{
		{"zero--", 0, 0},
		{"one--", 1, 0},
		{"large--", 2021, 2020},
	}

	for _, v := range scenario {
		input := Streak(v.input)
		input.Decrement()
		output := Streak(v.output)

		assert.Equal(t, output, input, v.desc)
	}

}
