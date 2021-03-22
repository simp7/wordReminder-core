package voca

type Meanings struct {
	Data []Meaning
}

func (m Meanings) IsEqual(u Unit) bool {

	ok := false

	for _, v := range m.Data {
		ok = ok || u.IsEqual(v)
	}

	return ok

}
