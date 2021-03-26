package spelling

import (
	"github.com/simp7/wordReminder-core/voca/unit"
	"testing"
)

func TestSpelling_IsRight(t *testing.T) {

	scenario := []struct {
		desc     string
		spelling unit.Spelling
		ans      string
		output   bool
	}{
		{"apple", New("apple"), "apple", true},
		{"apple", New("apple"), "banana", false},
		{"go", New("go"), "go", true},
		{"trial and error", New("trial and error"), "trials", false},
	}

	for _, v := range scenario {
		if v.spelling.IsRight(v.ans) == v.output {
			t.Logf("test \"%s\" has been passed!\n", v.desc)
		} else {
			t.Errorf("error in test \"%s\": wanted %t, but got %t\n", v.desc, v.output, v.spelling.IsRight(v.ans))
		}
	}

}

func TestSpelling_String(t *testing.T) {

	scenario := []struct {
		desc     string
		spelling unit.Spelling
		output   string
	}{
		{"apple", New("apple"), "apple"},
		{"go", New("go"), "go"},
		{"beautiful", New("beautiful"), "beautiful"},
	}

	for _, v := range scenario {
		if v.spelling.String() == v.output {
			t.Logf("test \"%s\" has been passed!\n", v.desc)
		} else {
			t.Errorf("error in test \"%s\": wanted %s, but got %s\n", v.desc, v.output, v.spelling)
		}
	}

}
