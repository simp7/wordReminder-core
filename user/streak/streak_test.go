package streak

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	scenario := []struct {
		desc    string
		input   int
		isError bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"negative", -1, true},
	}

	for _, v := range scenario {
		_, err := New(v.input)
		if v.isError {
			assert.Error(t, err, v.desc)
		} else {
			assert.NoError(t, err, v.desc)
		}
	}

}

func TestStreak_Increment(t *testing.T) {

	scenario := []struct {
		desc   string
		input  int
		output int
	}{
		{"zero++", 0, 1},
		{"one++", 1, 2},
	}

	for _, v := range scenario {
		input, _ := New(v.input)
		input.Increment()
		output, _ := New(v.output)

		assert.Equal(t, output, input, v.desc)
	}

}
