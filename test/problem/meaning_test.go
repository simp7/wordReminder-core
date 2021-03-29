package problem

import (
	"github.com/simp7/wordReminder-core/test"
	"github.com/simp7/wordReminder-core/voca/meaning"
	"github.com/simp7/wordReminder-core/voca/word"
	"testing"
)

func TestProblem_IsCorrect(t *testing.T) {

	problems := GetProblems()
	input := []string{"사과", "apple", "시험하다", "시험하다", "zebra", "go", "do", "가다", "가다"}
	mark := []bool{true, true, true, false, true, true, false, false, true}

	if len(problems) != len(input) || len(problems) != len(mark) {
		t.Errorf("Length of slice is different")
		return
	}

	for i, v := range problems {
		Evaluate(t, i, v.IsCorrect(input[i]), mark[i])
	}
}

func TestProblem_Question(t *testing.T) {

	problems := GetProblems()
	ans := []string{"apple", "사과", "test", "exam", "얼룩말", "가다, Go 언어", "가다, Go 언어", "가다, Go 언어", "go"}

	if len(problems) != len(ans) {
		t.Errorf("Length of slice is different")
		return
	}

	for i, v := range problems {
		Evaluate(t, i, v.Question().String(), ans[i])
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

	a := word.New("apple", meaning.Noun("사과"))
	t := word.New("test", meaning.Verb("시험하다"), meaning.Noun("시험"))
	e := word.New("exam", meaning.Noun("시험"))
	z := word.New("zebra", meaning.Noun("얼룩말"))
	g := word.New("go", meaning.Verb("가다"), meaning.Noun("Go 언어"))

	p1 := Meaning(a)
	p2 := Spelling(a)
	p3 := Meaning(t)
	p4 := Meaning(e)
	p5 := Spelling(z)
	p6 := Spelling(g)
	p7 := Spelling(g)
	p8 := Spelling(g)
	p9 := Meaning(g)

	return []test.Problem{p1, p2, p3, p4, p5, p6, p7, p8, p9}

}
