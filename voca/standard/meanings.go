package standard

import (
	"github.com/simp7/wordReminder-core/voca"
)

type meanings struct {
	Elements []voca.Meaning
}

func Meanings(mean ...voca.Meaning) voca.Meanings {
	return meanings{Elements: mean}
}

func (m meanings) IsRight(input string) bool {

	ok := false

	for _, v := range m.Elements {
		ok = ok || v.IsRight(input)
	}

	return ok

}

func (m meanings) String() string {

	var result string

	for i, v := range m.Elements {
		result += v.String()
		if len(m.Elements) > i+1 {
			result += ", "
		}
	}

	return result

}
