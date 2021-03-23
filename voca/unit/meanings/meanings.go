package meanings

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/simp7/wordReminder-core/voca/unit"
)

type meanings struct {
	Data []unit.Meaning
}

func New(mean ...unit.Meaning) unit.Meanings {
	return meanings{Data: mean}
}

func (m meanings) IsEqual(u voca.Unit) bool {

	ok := false

	for _, v := range m.Data {
		ok = ok || u.IsEqual(v)
	}

	return ok

}

func (m meanings) Format() string {

	var result string

	for i, v := range m.Data {
		result += v.Format()
		if len(m.Data) > i+1 {
			result += ", "
		}
	}

	return result

}
