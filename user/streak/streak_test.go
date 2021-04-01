package streak

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	scenario := []struct {
		desc   string
		input  int
		output error
	}{
		{"zero", 0, nil},
		{"one", 1, nil},
		{"negative", -1, NegativeStreakErr},
	}

	for _, v := range scenario {
		_, err := New(v.input)
		assert.Equal(t, v.output, err, v.desc)
	}

}

func TestStreak_String(t *testing.T) {

	scenario := []struct {
		desc   string
		input  int
		output string
	}{
		{"zero", 0, "0"},
		{"one", 1, "1"},
	}

	for _, v := range scenario {
		st, _ := New(v.input)
		assert.Equal(t, v.output, st.String())
	}

}

func TestStreak_Increment(t *testing.T) {

	scenario := []struct {
		desc   string
		input  int
		output string
	}{
		{"zero++", 0, "1"},
		{"one++", 1, "2"},
	}

	for _, v := range scenario {
		st, _ := New(v.input)
		st.Increment()
		assert.Equal(t, v.output, st.String(), v.desc)
	}

}
