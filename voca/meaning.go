package voca

type Meaning struct {
	Class Class
	Mean  string
}

func (m Meaning) IsEqual(u Unit) bool {
	return m == u
}
