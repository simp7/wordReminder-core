package voca

type WordSet interface {
	Len() int
	Get(idx int) Word
	Less(int, int) bool
	Swap(int, int)
}