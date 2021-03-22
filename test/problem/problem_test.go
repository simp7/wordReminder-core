package problem

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/test/input"
	"github.com/simp7/wordReminder-core/test/problemType"
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/class"
	"github.com/simp7/wordReminder-core/voca/meaning"
	"github.com/simp7/wordReminder-core/voca/meanings"
	"github.com/simp7/wordReminder-core/voca/spelling"
	"github.com/simp7/wordReminder-core/voca/word"
	"testing"
)

func TestProblem_IsCorrect(t *testing.T) {

	problems := GetProblems()
	ans := []voca.Unit{}
	mark := []bool{true, false}

	if len(problems) != len(ans) || len(problems) != len(mark) {
		t.Errorf("Length of slice is different")
		return
	}

	for i, v := range problems {
		Evaluate(t, i, v.IsCorrect(ans[i]), mark)
	}

}

func TestProblem_Question(t *testing.T) {

	problems := GetProblems()
	ans := []voca.Unit{spelling.New("apple"),
		meaning.New(class.Noun, "사과"),
		spelling.New("test"),
		spelling.New("exam"),
		meaning.New(class.Noun, "얼룩말"),
		meaning.New(class.Noun, "Go 언어"),
		meaning.New(class.Verb, "가다"),
		meanings.New(meaning.New(class.Verb, "오다"), meaning.New(class.Noun, "Go 언어")), // Should return error
		spelling.New("go"),
	}

	if len(problems) != len(ans) {
		t.Errorf("Length of slice is different")
		return
	}

	for i, v := range problems {
		if !(v.Question() == ans[i]) {
			t.Errorf("Errors in %d : wanted %+v, but got %+v", i, v.Question(), ans[i])
		} else {
			t.Logf("Test %d has been passed!", i)
		}
	}

}

func Evaluate(t *testing.T, idx int, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Errors in %d : wanted %+v, but got %+v", idx, a, b)
	} else {
		t.Logf("Test %d has been passed!", idx)
	}
}

func GetProblems() []test.Problem {

	a := word.New("apple", meaning.New(class.Noun, "사과"))
	t := word.New("test", meaning.New(class.Verb, "시험하다"), meaning.New(class.Noun, "시험"))
	e := word.New("exam", meaning.New(class.Verb, "시험하다"), meaning.New(class.Noun, "시험"))
	z := word.New("zebra", meaning.New(class.Noun, "얼룩말"))
	g := word.New("go", meaning.New(class.Verb, "가다"), meaning.New(class.Noun, "Go 언어"))

	p1 := New(a, input.Short(), problemType.Meaning())
	p2 := New(a, input.Short(), problemType.Spelling())
	p3 := New(t, input.Short(), problemType.Meaning())
	p4 := New(e, input.Short(), problemType.Meaning())
	p5 := New(z, input.Choice(4), problemType.Spelling())
	p6 := New(g, input.Short(), problemType.Spelling())
	p7 := New(g, input.Choice(4), problemType.Spelling())
	p8 := New(g, input.Choice(5), problemType.Spelling())
	p9 := New(g, input.Choice(4), problemType.Meaning())

	return []test.Problem{p1, p2, p3, p4, p5, p6, p7, p8, p9}

}
