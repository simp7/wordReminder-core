package meaning

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/class"
	"testing"
)

func TestMeaning_IsRight(t *testing.T) {

	scenario := []struct {
		desc    string
		meaning voca.Meaning
		ans     string
		output  bool
	}{
		{"apple", Noun("사과"), "사과", true},
		{"go", Verb("가다"), "가다", true},
		{"trial and error", Idiom("시행착오"), "실패", false},
	}

	for _, v := range scenario {
		if v.meaning.IsRight(v.ans) == v.output {
			t.Logf("test \"%s\" has been passed!\n", v.desc)
		} else {
			t.Errorf("error in test \"%s\": wanted %t, but got %t\n", v.desc, v.output, v.meaning.IsRight(v.ans))
		}
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
		if v.meaning.String() == v.output {
			t.Logf("test \"%s\" has been passed!\n", v.desc)
		} else {
			t.Errorf("error in test \"%s\": wanted %s, but got %s\n", v.desc, v.output, v.meaning)
		}
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
		if v.meaning.Type() == v.output {
			t.Logf("test \"%s\" has been passed!\n", v.desc)
		} else {
			t.Errorf("error in test \"%s\": wanted %d, but got %d\n", v.desc, v.output, v.meaning.Type())
		}
	}

}
