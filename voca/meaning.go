package voca

type Meaning interface {
	Unit
	Type() Class
}
