package problems

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/test/input"
	"github.com/simp7/wordReminder-core/test/problem"
	"github.com/simp7/wordReminder-core/voca/meaning"
	"github.com/simp7/wordReminder-core/voca/word"
	"testing"
)

func TestArray_Iterate(t *testing.T) {
	scenario := []struct {
		desc     string
		problems test.Problems
		input    []string
	}{
		{"apple", Array(problem.Meaning(word.New("apple", meaning.Noun("사과")), input.Short())), []string{"사과"}},
	}

	for _, v := range scenario {
		i := 0
		ok := true
		v.problems.Iterate(func(problem test.Problem) {
			if !problem.IsCorrect(v.input[i]) {
				ok = false
				return
			}
			i += 1
		})
		if ok {
			t.Logf("test \"%s\" has been passed!\n", v.desc)
		} else {
			t.Errorf("error in test \"%s\"\n ", v.desc)
		}
	}
}
