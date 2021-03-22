package voca

type Spelling struct {
	Data string
}

func (s Spelling) IsEqual(u Unit) bool {
	return s == u
}
