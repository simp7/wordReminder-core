package voca

import "sort"

type WordSet interface {
	sort.Interface
	Get(idx int) Word
}
