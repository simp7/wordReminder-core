package meanings

import (
	"github.com/simp7/wordReminder-core/voca/class"
	"github.com/simp7/wordReminder-core/voca/unit"
	"github.com/simp7/wordReminder-core/voca/unit/meaning"
	"testing"
)

func TestMeanings_IsRight(t *testing.T) {

	scenario := []struct {
		desc     string
		meanings unit.Meanings
		ans      string
		output   bool
	}{
		{"apple", New(meaning.New(class.Noun, "사과")), "사과", true},
		{"apple", New(meaning.New(class.Noun, "사과")), "바나나", false},
		{"go", New(meaning.New(class.Verb, "가다"), meaning.New(class.Noun, "Go 언어")), "가다", true},
		{"go", New(meaning.New(class.Verb, "가다"), meaning.New(class.Noun, "Go 언어")), "Go 언어", true},
		{"trial and error", meaning.New(class.Idiom, "시행착오"), "실패", false},
	}

	for _, v := range scenario {
		if v.meanings.IsRight(v.ans) == v.output {
			t.Logf("test \"%s\" has been passed!\n", v.desc)
		} else {
			t.Errorf("error in test \"%s\": wanted %t, but got %t\n", v.desc, v.output, v.meanings.IsRight(v.ans))
		}
	}

}

func TestMeanings_String(t *testing.T) {

	scenario := []struct {
		desc     string
		meanings unit.Meanings
		output   string
	}{
		{"apple", New(meaning.New(class.Noun, "사과")), "사과"},
		{"go", New(meaning.New(class.Verb, "가다"), meaning.New(class.Noun, "Go 언어")), "가다, Go 언어"},
		{"beautiful", New(meaning.New(class.Adjective, "아름다운")), "아름다운"},
	}

	for _, v := range scenario {
		if v.meanings.String() == v.output {
			t.Logf("test \"%s\" has been passed!\n", v.desc)
		} else {
			t.Errorf("error in test \"%s\": wanted %s, but got %s\n", v.desc, v.output, v.meanings)
		}
	}

}
